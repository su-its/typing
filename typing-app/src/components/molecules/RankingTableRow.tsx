import { Td, Tr } from "@chakra-ui/react";
import { components } from "@/libs/api/v0";

const RankingTableRow: React.FC<components["schemas"]["ScoreRanking"]> = (scoreRanking) => {
  const accuracy = scoreRanking.score?.accuracy ?? 0;

  let formattedAccuracy: string;

  let accuracyPercent = (accuracy * 100).toFixed(2);

  if (accuracyPercent === "100.00") {
    formattedAccuracy = "100%";
  } else if (accuracyPercent.endsWith(".00")) {
    formattedAccuracy = accuracyPercent.slice(0, -3) + "%";
  } else if (accuracyPercent.endsWith("0")) {
    formattedAccuracy = accuracyPercent.slice(0, -1) + "%";
  } else {
    formattedAccuracy = accuracyPercent + "%";
  }

  const formattedCreatedAt = scoreRanking.score?.created_at
    ? new Date(scoreRanking.score.created_at).toISOString().split("T")[0]
    : "";

  return (
    <Tr _even={{ bg: "midnightblue" }} _odd={{ bg: "#192f70" }} color={"silver"}>
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
        {formattedAccuracy}
      </Td>
      <Td width={"320px"} textAlign={"center"}>
        {formattedCreatedAt}
      </Td>
    </Tr>
  );
};

export default RankingTableRow;
