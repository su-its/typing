import { Tbody } from "@chakra-ui/react";
import RankingTableRow from "./RankingTableRow";
import { components } from "@/libs/api/v1";

export type RankingTableBodyProps = {
  scoreRankings: components["schemas"]["ScoreRanking"][];
};

const RankingTableBody: React.FC<RankingTableBodyProps> = ({ scoreRankings }) => {
  return (
    <Tbody>
      {scoreRankings.map((scoreRanking) => (
        <RankingTableRow key={String(scoreRanking.score?.user?.student_number)} {...scoreRanking} />
      ))}
    </Tbody>
  );
};
export default RankingTableBody;
