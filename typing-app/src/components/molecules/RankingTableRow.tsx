import { Td, Tr } from "@chakra-ui/react";
import { components } from "@/libs/api/v1";

const RankingTableRow: React.FC<components["schemas"]["ScoreRanking"]> = (scoreRanking) => {
  return (
    <Tr
      key={String(scoreRanking.score?.user?.student_number)}
      _even={{ bg: "midnightblue" }}
      _odd={{ bg: "#192f70" }}
      color={"silver"}
    >
      <Td width={"128px"} textAlign={"center"}>
        {String(scoreRanking.rank)}
      </Td>
      <Td width={"256px"} textAlign={"center"}>
        {scoreRanking.score?.user?.student_number}
      </Td>
      <Td width={"320px"} textAlign={"center"}>
        {String(scoreRanking.score?.keystrokes)}
      </Td>
      <Td width={"256px"} textAlign={"center"}>
        {String(scoreRanking.score?.accuracy)}
      </Td>
      <Td width={"320px"} textAlign={"center"}>
        {scoreRanking.score?.created_at}
      </Td>
    </Tr>
  );
};

export default RankingTableRow;
