import styles from "@/assets/sass/molecules/RankingTableHead.module.scss";

const RankingTableHead: React.FC = () => {
  return (
    <thead className={styles.head}>
      <tr>
        <th>順位</th>
        <th>学籍番号</th>
        <th>名前</th>
        <th>入力文字数</th>
        <th>正打率</th>
        <th>記録日</th>
      </tr>
    </thead>
  );
};

export default RankingTableHead;
