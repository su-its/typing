import styles from "@/assets/sass/molecules/RankingTableRow.module.scss";
import type { ColumnDefinition } from "../organism/RankingTabs";

interface EmptyTableRowProps {
  columns: ColumnDefinition[];
}

const EmptyTableRow: React.FC<EmptyTableRowProps> = ({ columns }) => {
  return (
    <tr className={styles.row}>
      {columns.map((column, index) => (
        <td key={column.key} className={index === 0 ? styles.rank : undefined}>
          -
        </td>
      ))}
    </tr>
  );
};

export default EmptyTableRow;
