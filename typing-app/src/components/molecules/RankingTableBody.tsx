import RankingTableRow from "./RankingTableRow";
import { components } from "@/libs/api/v0";

export type RankingTableBodyProps = {
  scoreRankings: components["schemas"]["ScoreRanking"][];
  displayRows: number;
};

const RankingTableBody: React.FC<RankingTableBodyProps> = ({ scoreRankings, displayRows }) => {
  return (
    <tbody>
      {Array.from({ length: displayRows }).map((_, index) => {
        const scoreRanking = scoreRankings[index];
        const rank = index + 1;

        // データが存在する場合はそのデータを、存在しない場合はrankのみを持つオブジェクトを渡す
        const rowData = scoreRanking ? scoreRanking : { rank: rank, score: undefined }; // score は undefined を渡す

        // キーはデータがある場合は score.id、ない場合は rank を使う
        const key = scoreRanking?.score?.id ?? `rank-${rank}`;

        return <RankingTableRow key={key} {...rowData} />;
      })}
    </tbody>
  );
};

export default RankingTableBody;
