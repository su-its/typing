import { components } from "@/libs/api/v0";
import styles from "@/assets/sass/molecules/RankingTableRow.module.scss";

const RankingTableRow: React.FC<components["schemas"]["ScoreRanking"]> = (scoreRanking) => {
  const accuracy = scoreRanking.score?.accuracy ?? 0;

  const formatter = new Intl.NumberFormat("en-US", {
    style: "percent",
    maximumFractionDigits: 2,
  });

  const formattedAccuracy = formatter.format(accuracy);

  const formattedCreatedAt = scoreRanking.score?.created_at
    ? new Date(scoreRanking.score.created_at).toISOString().split("T")[0]
    : "";

  return (
    <tr className={styles.row}>
      <td className={styles.rank}>{String(scoreRanking.rank)}</td>
      <td>{scoreRanking.score?.user?.student_number}</td>
      <td>{String(scoreRanking.score?.keystrokes)}</td>
      <td>{formattedAccuracy}</td>
      <td>{formattedCreatedAt}</td>
    </tr>
  );
};

export default RankingTableRow;
