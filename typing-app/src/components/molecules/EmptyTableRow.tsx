import styles from "@/assets/sass/molecules/RankingTableRow.module.scss"; // スタイルを再利用

const EmptyTableRow: React.FC = () => {
	// 表示する列の数を RankingTableRow と合わせる
	const numberOfColumns = 5; // ランク以外のデータ列の数

	return (
		<tr className={styles.row}>
			<td className={styles.rank}>{String("-")}</td>
			{/* データがない列は "-" を表示 */}
			{Array.from({ length: numberOfColumns }).map((_, columnIndex) => (
				<td key={`empty-cell-column-${columnIndex}`}>-</td>
			))}
		</tr>
	);
};

export default EmptyTableRow;
