import { Td, Tr } from "@chakra-ui/react";
import { ScoreRanking } from "../organism/RankingTabs";

const RankingTableRow: React.FC<ScoreRanking> = (scoreRanking) => {
  return (
    <Tr key={String(scoreRanking.rank)} _even={{bg: 'midnightblue'}} _odd={{bg: '#192f70'}} color={'silver'}>
      <Td width={"128px"} textAlign={"center"}>{String(scoreRanking.rank)}</Td>
      <Td isNumeric width={"256px"} textAlign={"center"}>{scoreRanking.user.studentNumber}</Td>
      <Td width={"320px"} textAlign={"center"}>{String(scoreRanking.keystrokes)}</Td>
      <Td width={"256px"} textAlign={"center"}>{String(scoreRanking.accuracy)}</Td>
      <Td width={"320px"} textAlign={"center"}>{scoreRanking.createdAt.toLocaleDateString("ja-JP")}</Td>
    </Tr>
  );
};

export default RankingTableRow;
