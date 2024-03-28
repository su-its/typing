import { Table, TableContainer, Container } from "@chakra-ui/react";
import RankingTableHead from "../molecules/RankingTableHead";
import RankingTableBody, { RankingTableBodyProps } from "../molecules/RankingTableBody";

const RankingTable: React.FC<RankingTableBodyProps> = ({ scoreRankings }) => {
  return (
    <TableContainer>
      <Container>
        <Table variant='striped' bg='white' rounded='base' shadow='md'>
          <RankingTableHead />
          <RankingTableBody scoreRankings={scoreRankings} />
        </Table>
      </Container>
    </TableContainer>
  );
};

export default RankingTable;
