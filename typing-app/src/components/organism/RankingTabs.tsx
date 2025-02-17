"use client";

import { Tabs, TabList, TabPanels, Tab, TabPanel, Flex, Center, Box, Grid } from "@chakra-ui/react";
import RankingTable from "../organism/RankingTable";
import { Pagination } from "../molecules/Pagination";
import RefreshButton from "../atoms/RefreshButton";
import { useCallback, useEffect, useState } from "react";
import { client } from "@/libs/api";
import { components } from "@/libs/api/v0";
import { showErrorToast } from "@/utils/toast";

const ITEMS_PER_PAGE = 10;

export default function RankingTabs() {
  const [data, setData] = useState<{
    rankings: components["schemas"]["ScoreRanking"][];
    totalCount: number;
  }>({
    rankings: [],
    totalCount: 0,
  });
  const [isLoading, setIsLoading] = useState(true);
  const [currentPage, setCurrentPage] = useState(1);
  const [sortBy, setSortBy] = useState<"accuracy" | "keystrokes">("accuracy");

  const fetchRankingData = useCallback(async () => {
    try {
      setIsLoading(true);
      const { data: responseData } = await client.GET("/scores/ranking", {
        params: {
          query: {
            sort_by: sortBy,
            start: (currentPage - 1) * ITEMS_PER_PAGE + 1,
            limit: ITEMS_PER_PAGE,
          },
        },
      });

      if (responseData) {
        setData({
          rankings: responseData.rankings,
          totalCount: responseData.total_count,
        });
      } else {
        throw new Error("ランキングデータの取得に失敗しました");
      }
    } catch (error) {
      showErrorToast("ランキングの取得に失敗しました");
      setData({ rankings: [], totalCount: 0 });
    } finally {
      setIsLoading(false);
    }
  }, [sortBy, currentPage]);

  useEffect(() => {
    fetchRankingData();
  }, [fetchRankingData]);

  const handleTabChange = (index: number) => {
    setSortBy(index === 0 ? "accuracy" : "keystrokes");
    setCurrentPage(1);
  };

  const handlePaginationClick = (direction: "next" | "prev") => {
    const newPage =
      direction === "prev"
        ? Math.max(currentPage - 1, 1)
        : Math.min(currentPage + 1, Math.ceil(data.totalCount / ITEMS_PER_PAGE));
    setCurrentPage(newPage);
  };

  return (
    <Tabs onChange={handleTabChange} mt={6}>
      <Flex justifyContent="center">
        <Grid templateColumns="repeat(3, 1fr)" gap="300px">
          <Box opacity="0" />
          <TabList color="white">
            <Tab _selected={{ color: "#00ace6" }}>正打率</Tab>
            <Tab _selected={{ color: "#00ace6" }}>入力文字数</Tab>
          </TabList>
          <RefreshButton onClick={fetchRankingData} isDisabled={isLoading} />
        </Grid>
      </Flex>

      <TabPanels>
        <TabPanel>
          <RankingTable scoreRankings={data.rankings} />
        </TabPanel>
        <TabPanel>
          <RankingTable scoreRankings={data.rankings} />
        </TabPanel>
      </TabPanels>

      <Center>
        <Pagination
          onPrev={() => handlePaginationClick("prev")}
          onNext={() => handlePaginationClick("next")}
          isPrevDisabled={currentPage <= 1}
          isNextDisabled={currentPage >= Math.ceil(data.totalCount / ITEMS_PER_PAGE)}
        />
      </Center>
    </Tabs>
  );
}
