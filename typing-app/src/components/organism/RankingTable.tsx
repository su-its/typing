import { Table, Thead, Tbody, Tfoot, Tr, Th, Td, TableCaption, TableContainer, Box } from "@chakra-ui/react";
import RankingTableHead from "../molecules/RankingTableHead";
import RankingTableBody from "../molecules/RankingTableBody";

interface User {
  id: string;
  studentNumber: string;
  handleName: string;
}
interface ScoreRanking {
  rank: Number;
  user: User;
  keystrokes: Number;
  accuracy: Number;
  createdAt: Date;
}

const demoUsers: User[] = [
  {
    id: "1",
    studentNumber: "X",
    handleName: "X",
  },
  {
    id: "2",
    studentNumber: "Y",
    handleName: "Y",
  },
  {
    id: "3",
    studentNumber: "Z",
    handleName: "Z",
  },
];

const demoScoreRankings: ScoreRanking[] = [
  {
    rank: 1,
    user: demoUsers[0],
    keystrokes: 100,
    accuracy: 100,
    createdAt: new Date(),
  },
  {
    rank: 2,
    user: demoUsers[1],
    keystrokes: 90,
    accuracy: 90,
    createdAt: new Date(),
  },
  {
    rank: 3,
    user: demoUsers[2],
    keystrokes: 80,
    accuracy: 80,
    createdAt: new Date(),
  },
  {
    rank: 4,
    user: demoUsers[2],
    keystrokes: 70,
    accuracy: 70,
    createdAt: new Date(),
  },
];

const RankingTable: React.FC = () => {
  return (
    <TableContainer>
      <Table colorScheme="black">
        <RankingTableHead />
        <RankingTableBody />
      </Table>
    </TableContainer>
  );
};

export default RankingTable;
