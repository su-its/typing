"use client";

import RankingTable from "../organism/RankingTable";
import { Pagination } from "../molecules/Pagination";
import RefreshButton from "../atoms/RefreshButton";
import { useEffect, useState } from "react";
import { client } from "@/libs/api";
import { components } from "@/libs/api/v0";
import { showErrorToast } from "@/utils/toast";
import styles from "@/assets/sass/organism/RankingTabs.module.scss";

const RankingTabs = () => {
  const [scoreRankings, setScoreRankings] = useState<components["schemas"]["ScoreRanking"][]>([]);
  const [rankingStartFrom, setRankingStartFrom] = useState(1);
  const [sortBy, setSortBy] = useState<"accuracy" | "keystrokes">("accuracy");
  const [totalRankingCount, setTotalRankingCount] = useState<number>(0);

  const LIMIT = 10; //TODO: Configファイルから取得

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
            {(() => {
              if (sortBy === "accuracy") {
                return (
                  <>
                    <div className={`${styles.tab} ${styles.selected}`} onClick={() => handleTabChange(0)}>
                      正打率
                    </div>
                    <div className={styles.tab} onClick={() => handleTabChange(1)}>
                      入力文字数
                    </div>
                  </>
                );
              } else {
                return (
                  <>
                    <div className={styles.tab} onClick={() => handleTabChange(0)}>
                      正打率
                    </div>
                    <div className={`${styles.tab} ${styles.selected}`} onClick={() => handleTabChange(1)}>
                      入力文字数
                    </div>
                  </>
                );
              }
            })()}
          </div>
          <RefreshButton
            onClick={() => {
              setRankingStartFrom(1);
              fetchData();
            }}
            isDisabled={false}
          />
        </div>
        <RankingTable scoreRankings={scoreRankings} />
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
