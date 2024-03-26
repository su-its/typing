interface RegisterScore {
  Keystrokes: number;
  Accuracy: number;
  Score: number;
  StartedAt: Date;
  EndedAt: Date;
}

export interface ResultScore {
  Keystrokes: number;
  Miss: number;
  Time: Date;
  WPM: number;
  Accuracy: number;
}

export default RegisterScore;
