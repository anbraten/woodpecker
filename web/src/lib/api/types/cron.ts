export interface Cron {
  id: number;
  name: string;
  branch: string;
  schedule: string;
  next_exec: number;
}
