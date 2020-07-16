export default function ProgressRound(){
    return (
                        
    <div className="container mb-5">
        <h1 className="col-lg-8 bold-text mt-5 mb-5">My Most Used Languages</h1>
        <div className="row">
            
            <div className="col-md-3 col-sm-6 pb-5">
                <div className="progress-round red">
                    <span className="progress-left">
                        <span className="progress-round-bar"></span>
                    </span>
                    <span className="progress-right">
                        <span className="progress-round-bar"></span>
                    </span>
                    <div className="progress-round-value">[69%]</div>
                </div>
                <h4 className="pt-4 text-center bold-text">Lang1</h4>
            </div>

            <div className="col-md-3 col-sm-6 pb-5">
                <div className="progress-round yellow">
                    <span className="progress-left">
                        <span className="progress-round-bar"></span>
                    </span>
                    <span className="progress-right">
                        <span className="progress-round-bar"></span>
                    </span>
                    <div className="progress-round-value">[75%]</div>
                </div>
                <h4 className="pt-4 text-center bold-text">Lang2</h4>
            </div>
            
            <div className="col-md-3 col-sm-6 pb-5">
                <div className="progress-round blue">
                    <span className="progress-left">
                        <span className="progress-round-bar"></span>
                    </span>
                    <span className="progress-right">
                        <span className="progress-round-bar"></span>
                    </span>
                    <div className="progress-round-value">[100%]</div>
                </div>
                <h4 className="pt-4 text-center bold-text">Fun Had!</h4>
            </div>
        </div>
    </div>
    )
}