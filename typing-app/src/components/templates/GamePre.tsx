import { Box, Text } from "@chakra-ui/react";
import React, { useEffect } from "react";
import { SubGamePageProps } from "../pages/Game";

interface GamePreProps extends SubGamePageProps {
  filenames: string[];
  setSubjectTextData: (data: string) => void;
}

const GamePre: React.FC<GamePreProps> = ({ filenames, nextPage, setSubjectTextData }) => {

  // 0からnまでの乱数を生成する関数
  const getRandomInt = (maxNumber: number) => {
    return Math.floor(Math.random() * maxNumber);
  };

  // Spaceキーを押したときに実行する関数
  const handleSpaceButtonDown = async (e: KeyboardEvent) => {
    if (e.code === 'Space') {
      e.preventDefault();  // ページのスクロールなどのデフォルト動作を防止
      
      // ランダムにファイル名を選択
      const randomFile = filenames[getRandomInt(filenames.length)];
      // `public` ディレクトリからの相対パスを指定
      const filePath = `/texts/${randomFile}`;

      // fetch APIを使用してファイルの内容を読み込む
      try {
        const response = await fetch(filePath);
        const data = await response.text();
        setSubjectTextData(data); // 親コンポーネントや他のコンポーネントにデータを渡す場合
      } catch (error) {
        console.error('Error loading the text file:', error);
      }

      // 次のページへ
      nextPage();
    }
  };

  useEffect(() => {
    window.addEventListener('keydown', handleSpaceButtonDown);

    // コンポーネントのクリーンアップ時にイベントリスナーを削除
    return () => {
      window.removeEventListener('keydown', handleSpaceButtonDown);
    };
  }, []); // 空の依存配列を指定して、コンポーネントのマウント時にのみイベントリスナーを追加

  return (
    <Box>
      <Text>GamePre screen</Text>
      {/* <Button onKeyDown={handleSpaceButtonDown}>start</Button> */}
      <Text>{ filenames }</Text>
    </Box>
  );
};

export default GamePre;