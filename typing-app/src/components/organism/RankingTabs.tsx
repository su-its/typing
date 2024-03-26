"use client";
import { Tabs, TabList, TabPanels, Tab, TabPanel, Button, Flex } from "@chakra-ui/react";
import RankingTable from "../organism/RankingTable";
import { useEffect, useState } from "react";

const RankingTabs = () => {
  const [scoreRankings, setScoreRankings] = useState<ScoreRanking[]>([]);
  const [rankingStartFrom, setRankingStartFrom] = useState(0);
  const [sortBy, setSortBy] = useState<"accuracy" | "keystrokes">("accuracy");
  const LIMIT = 10;

  const MAXIMUM = 100; // TODO: MAXIMUMをAPIから取得する

  useEffect(() => {
    // ページが読み込まれたときにデータを取得
    fetchData();
  });

  const handleTabChange = (index: number) => {
    if (index === 0) {
      setSortBy("accuracy");
    } else if (index === 1) {
      setSortBy("keystrokes");
    }

    fetchData;
  };

  // 演算子を引数にとる、ボタンを押したときのハンドラ関数
  const handlePaginationClick = (direction: "next" | "prev") => {
    const newStartFrom =
      direction === "prev" ? Math.max(rankingStartFrom - LIMIT, 0) : Math.min(rankingStartFrom + LIMIT, MAXIMUM);
    setRankingStartFrom(newStartFrom);

    fetchData;
  };

  const fetchData = async () => {
    setScoreRankings(demoAccuracyRankings);
    // TODO: データをフェッチ
  };

  return (
    <Tabs onChange={handleTabChange}>
      <Flex justifyContent="right">
        <Button onClick={fetchData}>Reload</Button>
      </Flex>
      <TabList>
        <Tab>Accuracy</Tab>
        <Tab>KeyStoroke</Tab>
      </TabList>

      <TabPanels>
        <TabPanel>
          <RankingTable scoreRankings={demoAccuracyRankings} /> {/* TODO: scoreRankingsに置き換え */}
        </TabPanel>
        <TabPanel>
          <RankingTable scoreRankings={demoKeyStrokeRankings} /> {/* TODO: scoreRankingsに置き換え */}
        </TabPanel>
      </TabPanels>
      <Button onClick={() => handlePaginationClick("prev")} isDisabled={rankingStartFrom <= 0}>
        Prev
      </Button>
      <Button onClick={() => handlePaginationClick("next")} isDisabled={rankingStartFrom + LIMIT >= MAXIMUM}>
        {/* テストのときにMAXIMUM人が表示されているか確認する */}
        Next
      </Button>
    </Tabs>
  );
};
export default RankingTabs;

export interface User {
  id: string;
  studentNumber: string;
  handleName: string;
}

export interface ScoreRanking {
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

const demoKeyStrokeRankings: ScoreRanking[] = [
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

const demoAccuracyRankings: ScoreRanking[] = [
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
