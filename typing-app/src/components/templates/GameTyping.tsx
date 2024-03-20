import RegisterScore from "@/types/RegisterScore";
import { Box, Button, Progress, Text } from "@chakra-ui/react";
import axios from "axios";
import React, { useEffect, useState } from "react";
import { SubGamePageProps } from "../pages/Game";

const GameTyping: React.FC<SubGamePageProps> = ({ nextPage }) => {
  const totalSeconds = 20;
  const [count, setCount] = useState(totalSeconds);
  const damyScoreData = {
    Keystrokes: 123,
    Accuracy: 456.7,
    Score: 890.1,
    StartedAt: new Date(),
    EndedAt: new Date(),
  } as RegisterScore;
  const damyUserId = "damyId";

  const userId = damyUserId; // ToDo: 要変更
  const scoreData = damyScoreData; // ToDo: 要変更

  useEffect(() => {
    if (count <= 0) {
      axios
        .post(`http://localhost:8080/users/${userId}/scores`, scoreData)
        .then((res) => {
          console.log(res.data);
          nextPage();
        })
        .catch((error) => {
          console.error(error);
        });
    } else {
      const timer = setInterval(() => setCount(count - 1), 1000);
      return () => clearInterval(timer);
    }
  }, [count, nextPage]);

  const progress = ((totalSeconds - count) / totalSeconds) * 100;

  return (
    <Box>
      <Text>Typing screen</Text>
      <Button onClick={nextPage}>finish</Button>
      <Progress value={progress} colorScheme="blue" />
    </Box>
  );
};

export default GameTyping;
