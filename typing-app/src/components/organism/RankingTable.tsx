import RankingTableHead from "../molecules/RankingTableHead";
import RankingTableBody, { RankingTableBodyProps } from "../molecules/RankingTableBody";
import styles from "@/assets/sass/organism/RankingTable.module.scss";

const RankingTable: React.FC<RankingTableBodyProps> = ({ scoreRankings }) => {
  return (
    <div>
      <table className={styles.ranking}>
        <RankingTableHead />
        <RankingTableBody scoreRankings={scoreRankings} displayRows={10} />
      </table>
    </div>
  );
};

export default RankingTable;
