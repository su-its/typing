import { components } from "@/libs/api/v0";
import styles from "@/assets/sass/molecules/RankingTableRow.module.scss";

// 新しい Props 型を定義
type RankingTableRowProps = {
  rank: number;
  // score プロパティをオプショナルにする
  // score の型は ScoreRanking['score'] と同じにする
  score?: components["schemas"]["ScoreRanking"]["score"];
};

// コンポーネントの Props 型を新しい型に変更し、分割代入を使用
const RankingTableRow: React.FC<RankingTableRowProps> = ({ rank, score }) => {
  const accuracy = score?.accuracy ?? 0;

  const formatter = new Intl.NumberFormat("en-US", {
    style: "percent",
    maximumFractionDigits: 2,
  });

  const formattedAccuracy = formatter.format(accuracy);

  const formattedCreatedAt = score?.created_at ? new Date(score.created_at).toISOString().split("T")[0] : "";

  return (
    <tr className={styles.row}>
      <td className={styles.rank}>{String(rank)}</td>
      {/* score が存在しない場合は '-' を表示 */}
      <td>{score?.user?.student_number ?? "-"}</td>
      <td>{score?.user?.handle_name ?? "-"}</td>
      <td>{score?.keystrokes != null ? String(score.keystrokes) : "-"}</td>
      <td>{score ? formattedAccuracy : "-"}</td>
      <td>{score ? formattedCreatedAt : "-"}</td>
    </tr>
  );
};

export default RankingTableRow;
