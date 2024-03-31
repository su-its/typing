import { Thead, Tr, Th } from "@chakra-ui/react";

const RankingTableHead: React.FC = () => {
  return (
    <Thead>
      <Tr bg={"midnightblue"}>
        <Th color={"silver"}>順位</Th>
        <Th color={"silver"}>学籍番号</Th>
        <Th color={"silver"}>入力文字数</Th>
        <Th color={"silver"}>正打率</Th>
        <Th color={"silver"}>記録日</Th>
      </Tr>
    </Thead>
  );
};

export default RankingTableHead;
