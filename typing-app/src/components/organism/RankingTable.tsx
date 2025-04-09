import RankingTableHead from "../molecules/RankingTableHead";
import RankingTableBody from "../molecules/RankingTableBody";
import type { RankingTableBodyProps } from "../molecules/RankingTableBody";
import type { ColumnDefinition } from "./RankingTabs";
import styles from "@/assets/sass/organism/RankingTable.module.scss";

interface RankingTableProps extends RankingTableBodyProps {
  columns: ColumnDefinition[];
}

const RankingTable: React.FC<RankingTableProps> = ({ scoreRankings, displayRows, columns }) => {
  return (
    <div>
      <table className={styles.ranking}>
        <RankingTableHead columns={columns} />
        <RankingTableBody scoreRankings={scoreRankings} displayRows={displayRows} columns={columns} />
      </table>
    </div>
  );
};

export default RankingTable;
