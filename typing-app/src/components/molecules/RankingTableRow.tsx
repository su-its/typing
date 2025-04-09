import type { components } from "@/libs/api/v0";
import styles from "@/assets/sass/molecules/RankingTableRow.module.scss";
import type { ColumnDefinition } from "../organism/RankingTabs";

interface RankingTableRowProps {
  scoreRanking: components["schemas"]["ScoreRanking"];
  columns: ColumnDefinition[];
}

const RankingTableRow: React.FC<RankingTableRowProps> = ({ scoreRanking, columns }) => {
  return (
    <tr className={styles.row}>
      {/* columns に基づいてセルを動的にレンダリング */}
      {columns.map((column) => {
        // dataAccessor があればそれを使用し、なければキーでアクセス
        let cellData: React.ReactNode = null; // 型を ReactNode に
        if (column.key === "rank") {
          cellData = scoreRanking.rank;
        } else if (column.dataAccessor) {
          cellData = column.dataAccessor(scoreRanking);
        }

        // rank 列には特別なスタイルを適用
        const className = column.key === "rank" ? styles.rank : undefined;

        return (
          <td key={column.key} className={className}>
            {cellData}
          </td>
        );
      })}
    </tr>
  );
};

export default RankingTableRow;
