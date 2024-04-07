"use client";

import { Tabs, TabList, TabPanels, Tab, TabPanel, Flex, Center, Box, Grid } from "@chakra-ui/react";
import RankingTable from "../organism/RankingTable";
import { Pagination } from "../molecules/Pagination";
import RefreshButton from "../atoms/RefreshButton";
import { useEffect, useState } from "react";
import { client } from "@/libs/api";
import { components } from "@/libs/api/v1";
import { showErrorToast } from "@/utils/toast";

const RankingTabs = () => {
  const [scoreRankings, setScoreRankings] = useState<components["schemas"]["ScoreRanking"][]>([]);
  const [rankingStartFrom, setRankingStartFrom] = useState(1);
  const [sortBy, setSortBy] = useState<"accuracy" | "keystrokes">("accuracy");
  const [totalRankingCount, setTotalRankingCount] = useState<number>(0);

  const LIMIT = 10; //TODO: Configファイルから取得
  const MAXIMUM = totalRankingCount;

  const fetchData = async () => {
    const { data, error } = await client.GET("/scores/ranking", {
      params: {
        query: {
          sort_by: sortBy,
          start: rankingStartFrom,
          limit: LIMIT,
        },
      },
    });
    if (data) {
      setScoreRankings(data.rankings);
      setTotalRankingCount(data.total_count);
    } else {
      showErrorToast(error);
    }
  };
  useEffect(() => {
    fetchData();
  }, [sortBy, rankingStartFrom]);

  const handleTabChange = (index: number) => {
    const sortOption = index === 0 ? "accuracy" : "keystrokes";
    setSortBy(sortOption);
    setRankingStartFrom(1);
  };

  const handlePaginationClick = (direction: "next" | "prev") => {
    const newStartFrom =
      direction === "prev"
        ? Math.max(rankingStartFrom - LIMIT, 1)
        : Math.min(rankingStartFrom + LIMIT, MAXIMUM - LIMIT);
    setRankingStartFrom(newStartFrom);
  };

  return (
    <Tabs onChange={handleTabChange}>
      <Flex justifyContent={"center"}>
        <Grid templateColumns={"repeat(3, 1fr)"} gap={"300px"}>
          <Box opacity={"0"}>{/* 幅を揃えるためだけの要素，視覚的な意味はなし */}</Box>
          <TabList color={"white"}>
            <Tab _selected={{ color: "#00ace6" }}>正打率</Tab>
            <Tab _selected={{ color: "#00ace6" }}>入力文字数</Tab>
          </TabList>
          <RefreshButton
            onClick={() => {
              setRankingStartFrom(1);
              fetchData();
            }}
            isDisabled={false}
          />
        </Grid>
      </Flex>
      <TabPanels>
        <TabPanel>
          <RankingTable scoreRankings={scoreRankings} />
        </TabPanel>
        <TabPanel>
          <RankingTable scoreRankings={scoreRankings} />
        </TabPanel>
      </TabPanels>
      <Center>
        <Pagination
          onPrev={() => handlePaginationClick("prev")}
          onNext={() => handlePaginationClick("next")}
          isPrevDisabled={rankingStartFrom <= 1}
          isNextDisabled={rankingStartFrom + LIMIT >= MAXIMUM}
        />
      </Center>
    </Tabs>
  );
};

export default RankingTabs;
