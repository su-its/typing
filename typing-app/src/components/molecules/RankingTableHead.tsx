import { Thead, Tr, Th } from "@chakra-ui/react";

const RankingTableHead: React.FC = () => {
  return (
    <Thead>
      <Tr>
        <Th>Rank</Th>
        <Th>Student Number</Th>
        <Th>keyStrokes</Th>
        <Th>Accuracy</Th>
        <Th>Created at</Th>
      </Tr>
    </Thead>
  );
};

export default RankingTableHead;
