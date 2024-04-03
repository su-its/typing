import { Table, TableContainer, Container } from "@chakra-ui/react";
import RankingTableHead from "../molecules/RankingTableHead";
import RankingTableBody, { RankingTableBodyProps } from "../molecules/RankingTableBody";

const RankingTable: React.FC<RankingTableBodyProps> = ({ scoreRankings }) => {
  return (
    <TableContainer>
      <Container maxW={"container.xl"}>
        <Table rounded="base" shadow="md">
          <RankingTableHead />
          <RankingTableBody scoreRankings={scoreRankings} />
        </Table>
      </Container>
    </TableContainer>
  );
};

export default RankingTable;
