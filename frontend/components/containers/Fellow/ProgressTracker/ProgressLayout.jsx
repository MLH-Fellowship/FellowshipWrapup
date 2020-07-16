import ProgressTriangle from "./ProgressTriangle";
import ProgressRound from "./ProgressRound";

export default function ProgressLayout() {
  return (
    <div className="container">
      <div className="row">
        <h1 className="col-lg-12 display-3 text-left mt-5 section-heading">
          Progress compared to others in your team
        </h1>
        <h1 className="bg-text mb-2 pl-2">Progress Made</h1>

        <ProgressRound />
        <ProgressTriangle />
      </div>
    </div>
  );
}
