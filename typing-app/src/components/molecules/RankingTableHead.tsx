import { Thead, Tr, Th } from "@chakra-ui/react";

const RankingTableHead: React.FC = () => {
  return (
    <Thead>
      <Tr>
        <Th>Rank</Th>
        <Th>Student Number</Th>
        <Th isNumeric>keyStorokes</Th>
        <Th isNumeric>Accuracy</Th>
      </Tr>
    </Thead>
  );
};

export default RankingTableHead;
