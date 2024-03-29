import { Thead, Tr, Th } from "@chakra-ui/react";

const RankingTableHead: React.FC = () => {
  return (
    <Thead>
      <Tr bg={'midnightblue'}>
        <Th color={'silver'}>Rank</Th>
        <Th color={'silver'}>Student Number</Th>
        <Th color={'silver'}>keyStrokes</Th>
        <Th color={'silver'}>Accuracy</Th>
        <Th color={'silver'}>Created at</Th>
      </Tr>
    </Thead>
  );
};

export default RankingTableHead;
