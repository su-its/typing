import { Td, Tr } from "@chakra-ui/react";
import { ScoreRanking } from "../organism/RankingTabs";

const RankingTableRow: React.FC<ScoreRanking> = (scoreRanking) => {
  return (
    <Tr key={String(scoreRanking.rank)}>
      <Td>{String(scoreRanking.rank)}</Td>
      <Td>{scoreRanking.user.studentNumber}</Td>
      <Td isNumeric>{String(scoreRanking.keystrokes)}</Td>
      <Td isNumeric>{String(scoreRanking.accuracy)}</Td>
      <Td>{scoreRanking.createdAt.toLocaleDateString("ja-JP")}</Td>
    </Tr>
  );
};

export default RankingTableRow;
