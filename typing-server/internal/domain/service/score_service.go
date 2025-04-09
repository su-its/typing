package service

import (
	"errors"
	"sort"

	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/internal/domain/model"
	"github.com/su-its/typing/typing-server/internal/domain/repository"
)

// ScoreService はスコアのビジネスロジックを管理する
type ScoreService struct {
	scoreRepo repository.ScoreRepository
}

// NewScoreService は ScoreService のインスタンスを作成する
func NewScoreService(scoreRepo repository.ScoreRepository) *ScoreService {
	return &ScoreService{scoreRepo: scoreRepo}
}

// ValidateScore はスコアが有効かどうかを判定する
func (s *ScoreService) ValidateScore(userID uuid.UUID, keystrokes int, accuracy float64) error {
	if keystrokes < 0 {
		return errors.New("keystrokes must be non-negative")
	}
	if accuracy < 0 || accuracy > 1 {
		return errors.New("accuracy must be between 0 and 1")
	}
	if userID == uuid.Nil {
		return errors.New("invalid user ID")
	}
	return nil
}

// ComputeRanking は、スコア一覧を受け取り、ユーザーごとに一つのスコアに絞ったうえで
// sortByを基準にランキングを算出する例です。
// 返り値は「ランク付きスコア一覧」と「ランク対象件数」です。
func (s *ScoreService) ComputeRanking(scores []*model.Score, sortBy string) []*model.ScoreRanking {
	// 1. 同一ユーザーの重複スコアは無視して、ユーザーごとに最高スコアを1つだけ取得
	//    ここでは、既に .Order(ent_generated.Desc(sortBy)) でソート済みであることを想定して
	//    「最初に出てきたスコア ＝ 一番高い値のスコア」とみなしています。
	bestScoreByUser := make(map[string]*model.Score)
	for _, sc := range scores {
		if _, ok := bestScoreByUser[sc.UserID]; !ok {
			// userIDに紐づくスコアがまだなければ、これを登録
			bestScoreByUser[sc.UserID] = sc
		}
	}

	// 2. マップをスライスに変換
	var filteredScores []*model.Score
	for _, sc := range bestScoreByUser {
		filteredScores = append(filteredScores, sc)
	}

	// 3. 念のため sortBy で再度ソートしておく（上流でソート済みでも、ここで再ソートしておくと安全）
	//    sortBy = "accuracy" or "keystrokes" などを想定
	switch sortBy {
	case "accuracy":
		sort.Slice(filteredScores, func(i, j int) bool {
			return filteredScores[i].Accuracy > filteredScores[j].Accuracy
		})
	case "keystrokes":
		sort.Slice(filteredScores, func(i, j int) bool {
			return filteredScores[i].Keystrokes > filteredScores[j].Keystrokes
		})
	default:
		// sortByが他の値の場合のハンドリング（必要に応じて拡張）
		// ここではaccuracy降順をデフォルトにしておく
		sort.Slice(filteredScores, func(i, j int) bool {
			return filteredScores[i].Accuracy > filteredScores[j].Accuracy
		})
	}

	// 4. ランキングをつける
	//    - ここでは "RANK" のように「同値なら同じランク、次の順位は飛ばす」実装を例示
	//    - もし「DENSE_RANK」(同値なら同順位で、次は飛ばさない) にしたい場合はロジックを変更してください
	result := make([]*model.ScoreRanking, len(filteredScores))
	if len(filteredScores) == 0 {
		// 0件なら空のまま返す
		return result
	}

	currentRank := 1
	// 比較用の「前の要素のソート対象値」
	// sortByに応じて数値を取り出す。ここでは float64 に寄せて扱っています。
	getValue := func(sc *model.Score) float64 {
		switch sortBy {
		case "accuracy":
			return sc.Accuracy
		case "keystrokes":
			return float64(sc.Keystrokes)
		default:
			return sc.Accuracy // デフォルトはAccuracy
		}
	}

	lastValue := getValue(filteredScores[0])

	// 先頭要素はRank=1を付与
	result[0] = &model.ScoreRanking{
		Rank:  currentRank,
		Score: *filteredScores[0],
	}

	// 2番目以降
	for i := 1; i < len(filteredScores); i++ {
		thisValue := getValue(filteredScores[i])
		// 値が下がった(小さくなった)タイミングで次のランクへ
		if thisValue < lastValue {
			// 現在のインデックス i は0-basedなので、「i + 1」で1-basedの人数になる
			currentRank = i + 1
		}

		result[i] = &model.ScoreRanking{
			Rank:  currentRank,
			Score: *filteredScores[i],
		}
		lastValue = thisValue
	}

	return result
}

// LimitRankings はランキングを指定された範囲で安全に取得する
// startは1以上の値を想定
func (s *ScoreService) LimitRankings(rankings []*model.ScoreRanking, start, limit int) []*model.ScoreRanking {
	// 1から始まるインデックスを0ベースに変換
	zeroBasedStart := start - 1
	totalCount := len(rankings)

	// 範囲外の場合は空配列を返す
	if zeroBasedStart >= totalCount || zeroBasedStart < 0 {
		return []*model.ScoreRanking{}
	}

	// 終了位置の調整
	end := zeroBasedStart + limit
	if end > totalCount {
		end = totalCount
	}

	return rankings[zeroBasedStart:end]
}
