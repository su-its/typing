import { Thead, Tr, Th } from "@chakra-ui/react";

const RankingTableHead: React.FC = () => {
  return (
    <Thead>
      <Tr bg={"midnightblue"}>
        <Th color={"silver"} width={"128px"} textAlign={"center"}>順位</Th>
        <Th color={"silver"} width={"256px"} textAlign={"center"}>学籍番号</Th>
        <Th color={"silver"} width={"320px"} textAlign={"center"}>入力文字数</Th>
        <Th color={"silver"} width={"256px"} textAlign={"center"}>正打率</Th>
        <Th color={"silver"} width={"320px"} textAlign={"center"}>記録日</Th>
      </Tr>
    </Thead>
  );
};

export default RankingTableHead;
