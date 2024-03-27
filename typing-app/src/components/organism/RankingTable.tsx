import { Table, TableContainer, Box, Container } from "@chakra-ui/react";
import RankingTableHead from "../molecules/RankingTableHead";
import RankingTableBody, { RankingTableBodyProps } from "../molecules/RankingTableBody";

const RankingTable: React.FC<RankingTableBodyProps> = ({ scoreRankings }) => {
  return (
    <TableContainer>
      <Container>
      <Box bg="white" borderColor="black" rounded="base">
          <Table colorScheme="Black">
            <RankingTableHead />
            <RankingTableBody scoreRankings={scoreRankings} />
          </Table>
      </Box>
      </Container>
    </TableContainer>
  );
};

export default RankingTable;
