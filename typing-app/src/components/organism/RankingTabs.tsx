"use client";

import RankingTable from "../organism/RankingTable";
import { Pagination } from "../molecules/Pagination";
import RefreshButton from "../atoms/RefreshButton";
import { useCallback, useEffect, useState } from "react";
import { client } from "@/libs/api";
import type { components } from "@/libs/api/v0";
import { showErrorToast } from "@/utils/toast";
import styles from "@/assets/sass/organism/RankingTabs.module.scss";

// 列定義
const columns = [
  { key: "rank", label: "順位" },
  {
    key: "student_number",
    label: "学籍番号",
    dataAccessor: (scoreRanking: components["schemas"]["ScoreRanking"]) => scoreRanking.score.user.student_number,
  },
  {
    key: "handle_name",
    label: "ハンドルネーム",
    dataAccessor: (scoreRanking: components["schemas"]["ScoreRanking"]) => scoreRanking.score.user.handle_name,
  },
  {
    key: "keystrokes",
    label: "入力文字数",
    dataAccessor: (scoreRanking: components["schemas"]["ScoreRanking"]) => String(scoreRanking.score.keystrokes),
  },
  {
    key: "accuracy",
    label: "正打率",
    dataAccessor: (scoreRanking: components["schemas"]["ScoreRanking"]) => {
      const formatter = new Intl.NumberFormat("en-US", {
        style: "percent",
        maximumFractionDigits: 2,
      });
      return formatter.format(scoreRanking.score.accuracy);
    },
  },
  {
    key: "created_at",
    label: "記録日時",
    dataAccessor: (scoreRanking: components["schemas"]["ScoreRanking"]) =>
      new Date(scoreRanking.score.created_at).toISOString().split("T")[0],
  },
];

export type ColumnDefinition = (typeof columns)[number];

const RankingTabs = () => {
  const [scoreRankings, setScoreRankings] = useState<components["schemas"]["ScoreRanking"][]>([]);
  const [rankingStartFrom, setRankingStartFrom] = useState(1);
  const [sortBy, setSortBy] = useState<"accuracy" | "keystrokes">("accuracy");
  const [totalRankingCount, setTotalRankingCount] = useState<number>(0);

  const LIMIT = 10; //TODO: Configファイルから取得

  const fetchData = useCallback(async () => {
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
  }, [sortBy, rankingStartFrom]);

  useEffect(() => {
    fetchData();
  }, [fetchData]);

  const handleTabChange = (index: number) => {
    const sortOption = index === 0 ? "accuracy" : "keystrokes";
    setSortBy(sortOption);
    setRankingStartFrom(1);
  };

  const handlePaginationClick = (direction: "next" | "prev") => {
    const newStartFrom =
      direction === "prev"
        ? Math.max(rankingStartFrom - LIMIT, 1)
        : rankingStartFrom + LIMIT <= totalRankingCount
          ? rankingStartFrom + LIMIT
          : rankingStartFrom;
    setRankingStartFrom(newStartFrom);
  };

  return (
    <div className={styles.ranking}>
      <div className={styles.container}>
        <div className={styles.menu}>
          <div className={styles.tabs}>
            <button
              type="button"
              className={`${styles.tab} ${sortBy === "accuracy" ? styles.selected : ""}`}
              onClick={() => handleTabChange(0)}
            >
              正打率
            </button>
            <button
              type="button"
              className={`${styles.tab} ${sortBy !== "accuracy" ? styles.selected : ""}`}
              onClick={() => handleTabChange(1)}
            >
              入力文字数
            </button>
          </div>
          <div className={styles.refreshButtonWrapper}>
            <RefreshButton
              onClick={() => {
                setRankingStartFrom(1);
                fetchData();
              }}
            />
          </div>
        </div>
        <div className={styles.tableContainer}>
          <RankingTable scoreRankings={scoreRankings} displayRows={LIMIT} columns={columns} />
        </div>
        <div className={styles.pagination}>
          <Pagination
            onPrev={() => handlePaginationClick("prev")}
            onNext={() => handlePaginationClick("next")}
            isPrevDisabled={rankingStartFrom <= 1}
            isNextDisabled={rankingStartFrom + LIMIT > totalRankingCount}
          />
        </div>
      </div>
    </div>
  );
};

export default RankingTabs;
