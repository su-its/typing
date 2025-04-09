import RankingTableRow from "./RankingTableRow";
import EmptyTableRow from "./EmptyTableRow";
import type { components } from "@/libs/api/v0";

export type RankingTableBodyProps = {
	scoreRankings: components["schemas"]["ScoreRanking"][];
	displayRows: number;
};

const RankingTableBody: React.FC<RankingTableBodyProps> = ({
	scoreRankings,
	displayRows,
}) => {
	return (
		<tbody>
			{Array.from({ length: displayRows }).map((_, index) => {
				const scoreRanking = scoreRankings[index];
				if (scoreRanking) {
					return (
						<RankingTableRow key={scoreRanking.score.id} {...scoreRanking} />
					);
				}

				return <EmptyTableRow key={`rank-${index}`} />;
			})}
		</tbody>
	);
};

export default RankingTableBody;
