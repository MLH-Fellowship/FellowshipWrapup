import ProgressTriangle from "./ProgressTriangle";
import ProgressRound from "./ProgressRound";

export default function ProgressLayout({ issues, prs }) {
  return (
    <div className="container" style={{ marginBottom: "11.5rem" }}>
      <div className="row">
        <h1 className="col-lg-12 display-3 text-left mt-5 section-heading">
          Progress compared to others in your team
        </h1>
        <h1 className="bg-text mb-2 pl-2">Progress Made</h1>

        <ProgressRound />
        <ProgressTriangle issues={issues} prs={prs} />
      </div>
    </div>
  );
}
