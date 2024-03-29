import { Td, Tr } from "@chakra-ui/react";
import { ScoreRanking } from "../organism/RankingTabs";

const RankingTableRow: React.FC<ScoreRanking> = (scoreRanking) => {
  return (
    <Tr key={String(scoreRanking.rank)} _even={{bg: 'midnightblue'}} _odd={{bg: '#192f70'}} color={'silver'}>
      <Td>{String(scoreRanking.rank)}</Td>
      <Td isNumeric>{scoreRanking.user.studentNumber}</Td>
      <Td textAlign="center">{String(scoreRanking.keystrokes)}</Td>
      <Td textAlign="center">{String(scoreRanking.accuracy)}</Td>
      <Td>{scoreRanking.createdAt.toLocaleDateString("ja-JP")}</Td>
    </Tr>
  );
};

export default RankingTableRow;
