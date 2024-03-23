import { Td, Tr } from "@chakra-ui/react";
import { ScoreRanking } from "./RankingTableBody";

const RankingTableRow: React.FC<ScoreRanking> = (scoreRanking) => {
  return (
    <Tr key={String(scoreRanking.rank)}>
      <Td>{String(scoreRanking.rank)}</Td>
      <Td>{scoreRanking.user.studentNumber}</Td>
      <Td isNumeric>{String(scoreRanking.keystrokes)}</Td>
      <Td isNumeric>{String(scoreRanking.accuracy)}</Td>
    </Tr>
  );
};

export default RankingTableRow;
