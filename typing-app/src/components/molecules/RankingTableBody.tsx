import { Tbody } from "@chakra-ui/react";
import RankingTableRow from "./RankingTableRow";
import { components } from "@/libs/api/v0";

export type RankingTableBodyProps = {
  scoreRankings: components["schemas"]["ScoreRanking"][];
};

const RankingTableBody: React.FC<RankingTableBodyProps> = ({ scoreRankings }) => {
  return (
    <Tbody>
      {scoreRankings.map((scoreRanking) => (
        <RankingTableRow key={scoreRanking.score?.id ?? `rank-${scoreRanking.rank}`} {...scoreRanking} />
      ))}
    </Tbody>
  );
};

export default RankingTableBody;
