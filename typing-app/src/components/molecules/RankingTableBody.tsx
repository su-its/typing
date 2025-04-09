import RankingTableRow from "./RankingTableRow";
import EmptyTableRow from "./EmptyTableRow";
import type { components } from "@/libs/api/v0";
import type { ColumnDefinition } from "../organism/RankingTabs";

export type RankingTableBodyProps = {
  scoreRankings: components["schemas"]["ScoreRanking"][];
  displayRows: number;
  columns: ColumnDefinition[];
};

const RankingTableBody: React.FC<RankingTableBodyProps> = ({ scoreRankings, displayRows, columns }) => {
  return (
    <tbody>
      {Array.from({ length: displayRows }).map((_, index) => {
        const scoreRanking = scoreRankings[index];
        if (scoreRanking) {
          return <RankingTableRow key={scoreRanking.score.id} scoreRanking={scoreRanking} columns={columns} />;
        }

        return <EmptyTableRow key={`empty-row-${index}`} columns={columns} />;
      })}
    </tbody>
  );
};

export default RankingTableBody;
