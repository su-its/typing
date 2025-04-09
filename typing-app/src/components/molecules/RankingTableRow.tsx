import type { components } from "@/libs/api/v0";
import styles from "@/assets/sass/molecules/RankingTableRow.module.scss";

// 新しい Props 型を定義
type RankingTableRowProps = components["schemas"]["ScoreRanking"];

// コンポーネントの Props 型を新しい型に変更し、分割代入を使用
const RankingTableRow: React.FC<RankingTableRowProps> = ({ rank, score }) => {
	const accuracy = score.accuracy;

	const formatter = new Intl.NumberFormat("en-US", {
		style: "percent",
		maximumFractionDigits: 2,
	});

	const formattedAccuracy = formatter.format(accuracy);

	const formattedCreatedAt = new Date(score.created_at)
		.toISOString()
		.split("T")[0];

	return (
		<tr className={styles.row}>
			<td className={styles.rank}>{String(rank)}</td>
			<td>{score.user.student_number}</td>
			<td>{score.user.handle_name}</td>
			<td>{String(score.keystrokes)}</td>
			<td>{formattedAccuracy}</td>
			<td>{formattedCreatedAt}</td>
		</tr>
	);
};

export default RankingTableRow;
