import styles from "@/assets/sass/molecules/RankingTableHead.module.scss";
import type { ColumnDefinition } from "../organism/RankingTabs";

interface RankingTableHeadProps {
  columns: ColumnDefinition[];
}

const RankingTableHead: React.FC<RankingTableHeadProps> = ({ columns }) => {
  return (
    <thead className={styles.head}>
      <tr>
        {columns.map((column, index) => (
          <th key={column.key} className={index === 0 ? styles.rank : undefined}>
            {column.label}
          </th>
        ))}
      </tr>
    </thead>
  );
};

export default RankingTableHead;
