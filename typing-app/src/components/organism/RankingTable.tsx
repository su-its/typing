import { Table, TableContainer } from "@chakra-ui/react";
import RankingTableHead from "../molecules/RankingTableHead";
import RankingTableBody, { RankingTableBodyProps } from "../molecules/RankingTableBody";

const RankingTable: React.FC<RankingTableBodyProps> = ({ scoreRankings }) => {
  return (
    <TableContainer>
      <Table colorScheme="black">
        <RankingTableHead />
        <RankingTableBody scoreRankings={scoreRankings} />
      </Table>
    </TableContainer>
  );
};

export default RankingTable;
