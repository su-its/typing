import { Tbody } from "@chakra-ui/react";
import RankingTableRow from "./RankingTableRow";
import { ScoreRanking } from "../organism/RankingTabs";

export type RankingTableBodyProps = {
  scoreRankings: ScoreRanking[];
};

const RankingTableBody: React.FC<RankingTableBodyProps> = ({ scoreRankings }) => {
  return (
    <Tbody>
      {scoreRankings.map((scoreRanking) => (
        <RankingTableRow key={String(scoreRanking.rank)} {...scoreRanking} />
      ))}
    </Tbody>
  );
};
export default RankingTableBody;
