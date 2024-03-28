interface RegisterScore {
  keystrokes: number;
  accuracy: number;
  score: number;
  startedAt: Date;
  endedAt: Date;
}

export interface ResultScore {
  score: number;
  keystrokes: number;
  miss: number;
  time: Date;
  wpm: number;
  accuracy: number;
}

export default RegisterScore;
