import PodMember from "./PodMember";

export const PodInfo = ({ podInfo }) => {
    return (
      <StyledRepo>
            <div class="containter mt-5 mb-5">
                <div class="row">
                    <h1 class="col-lg-12 display-3 text-left mt-5 section-heading">The Team
                    </h1>
                    <h1 class="bold-text bg-text mb-2">[Pod x.x.x]</h1>
                    
                    <div class="row vertical-center pod-img">
                        <div class="col-lg-8 text-left pb-5">
                            <h1 class="display-1">[Full Name]</h1>
                        </div>
                        <div class="col-lg-4 text-center my-auto pb-5"><img class="img-fluid headerImg mb-3"
                                src="https://picsum.photos/512" /></div>
                    </div>
                    
                    <div class="container mb-5">
                        <div class="row justify-content-center">
                            <PodMember />
                        </div>
                    </div>
                </div>
            </div>
      </StyledRepo>
    );
  };
  
  export const StyledRepo = styled.div`
    .pod-img{
        min-height: 100% !important;
    }
  `;
  