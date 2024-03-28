import { Thead, Tr, Th } from "@chakra-ui/react";

const RankingTableHead: React.FC = () => {
  return (
    <Thead>
      <Tr>
        <Th>Rank</Th>
        <Th isNumeric>Student Number</Th>
        <Th isNumeric>keyStrokes</Th>
        <Th isNumeric>Accuracy</Th>
        <Th>Created at</Th>
      </Tr>
    </Thead>
  );
};

export default RankingTableHead;
