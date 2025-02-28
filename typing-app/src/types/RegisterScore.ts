interface RegisterScore {
  keystrokes: number;
  accuracy: number;
  score: number;
}
export interface ResultScore {
  score: number;
  keystrokes: number;
  miss: number;
  time: number;
  wpm: number;
  accuracy: number;
}

export default RegisterScore;
