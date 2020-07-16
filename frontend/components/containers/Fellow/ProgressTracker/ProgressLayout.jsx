import ProgressTriangle from "./ProgressTriangle"
import ProgressRound from "./ProgressRound";

export default function ProgressLayout() {
    return (
        <div class="container">
            <div class="row">
                <h1 class="col-lg-12 display-3 text-left mt-5 section-heading">Progress compared to others in
                    your team
                </h1>
                <h1 class="bg-text mb-2">Progress Made</h1>
                
                <ProgressRound />
                <ProgressTriangle />

            </div>
        </div>
    )
}