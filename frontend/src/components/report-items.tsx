import { type report } from "@/types/base/report";
import { formatDate } from "@/lib/utils";

type ReportItemProps = {
  report: report & {
    updatedAt: string;
  };
};

const ReportItem = ({ report }: ReportItemProps) => {
  return (
    <div className="flex items-center justify-between p-4">
      <div className="flex flex-col items-start gap-1">
        <div className="flex gap-2">
          <div className="font-semibold hover:underline">{report.topic}</div>
        </div>
        <p className="text-sm text-muted-foreground">
          {formatDate(parseInt(report.updatedAt) * 1000)}
        </p>
        <p className="text-md text-muted-foreground">{report.description}</p>
      </div>
    </div>
  );
};

export { ReportItem };
