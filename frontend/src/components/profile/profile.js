import React,{Component} from 'react';
import Header from "../Common/Header";
import axios from "axios";
import {Link} from "react-router-dom";
import {AUTH_API, USER_API} from "../constants/constants";


class Profile extends Component
{
    constructor(props)
    {
        super(props);

        this.state = {
            id : '',
            email: '',
            profile : '',
            firstname: '',
            lastname: '',
            skills: '',
            aboutme: '',
            company : '',
            role : '',
            roledesc : '',
            school : '',
            program : '',
            progdesc : ''
        }

        this.handleChange = this.handleChange.bind(this);
        this.submit = this.submit.bind(this);
    }

    componentDidMount()
    {
        console.log("In Did Mount success");
        var id = sessionStorage.getItem("user_id");
        axios.get(`${USER_API}/5cccde7d01aaf20001d743e6`).then((response)=>{
            console.log("In Get Start");
            if(response.data){
                this.setState({id : response.data.id});
                this.setState({email : response.data.email});
                this.setState({profile : response.data.profile});
                this.setState({firstname : response.data.firstname});
                this.setState({lastname : response.data.lastname});
                this.setState({skills : response.data.skills});
                this.setState({aboutme : response.data.aboutme});
                this.setState({company : response.data.experience.company});
                this.setState({role : response.data.experience.role});
                this.setState({roledesc : response.data.experience.desc});
                this.setState({school : response.data.education.school});
                this.setState({program : response.data.education.program});
                this.setState({progdesc : response.data.education.desc});
                console.log("In Get success");
            }
            console.log("In Get Failure");
        });

    }

    submit = (e) =>
    {
        e.preventDefault();

        const profiledata = {
            id : this.state.id,
            email: this.state.email,
            profile : this.state.profile,
            firstname: this.state.firstname,
            lastname: this.state.lastname,
            skills: this.state.skills,
            aboutme: this.state.aboutme,
            experience: {
                company : this.state.company,
                role: this.state.role,
                desc : this.state.roledesc
            },
            education: {
                school: this.state.school,
                program: this.state.program,
                desc : this.state.progdesc
            }
        }


        console.log("data sending:" + JSON.stringify(profiledata));
        axios.put(`${USER_API}/`, profiledata)
            .then((response) => {
                if (response.status === 202) {
                    alert('Profile update successfully');
                    this.componentDidMount();
                }

            });

    }



    handleChange = (e) =>
    {
        const target = e.target;
        const value = target.type === 'checkbox' ? target.checked : target.value;
        const name = target.name;

        this.setState({
            [name]: value
        });
    }
    render()
    {
        return(
            <div>
                <Header/>

                <div className="main-section pad-top-15">
                    <div className="container">
                        <div className="main-section-data">
                            <div className="row">

                                <div className="col-lg-9 col-md-8 no-pd">
                                    <div className="py-5">
                                        <div className="container">
                                            <div className="row">
                                                <div className="p-3 col-md-8">
                                                    <div className="blockquote mb-0">
                                                        <p>"Here comes the learning tag"</p>
                                                        <div className="blockquote-footer">
                                                            <b>J. W. Goethe</b>, CEO at Werther Inc.
                                                        </div>
                                                    </div>
                                                </div>
                                                <div className="col-md-2 mx-auto p-3">
                                                    <img
                                                    className="img-fluid d-block rounded-circle"
                                                    src="https://static.pingendo.com/img-placeholder-2.svg" />
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                        <hr></hr>

                                        <h1 className="mb-3">Profile</h1>
                                        <form className="needs-validation" noValidate="">
                                            <div className="row">
                                                <div className="col-md-6 mb-3">
                                                    <label htmlFor="firstName">First name</label>
                                                    <input type="text" className="form-control" name="firstname" placeholder="" value={this.state.firstname}
                                                           onChange={this.handleChange} />
                                                    <div className="invalid-feedback">
                                                        Valid first name is required.
                                                    </div>
                                                </div>
                                                <div className="col-md-6 mb-3">
                                                    <label htmlFor="lastName">Last name</label>
                                                    <input type="text" className="form-control" name="lastname" placeholder="" value={this.state.lastname}
                                                           onChange={this.handleChange} />
                                                    <div className="invalid-feedback">
                                                        Valid last name is required.
                                                    </div>
                                                </div>
                                            </div>

                                            <div className="mb-3">
                                                <label htmlFor="address">About Me</label>
                                                <input type="text" className="form-control" name="aboutme"
                                                       placeholder="Tell us something about you" onChange={this.handleChange} value={this.state.aboutme} />

                                            </div>

                                            <div className="mb-3">
                                                <label htmlFor="address">Skills</label>
                                                <input type="text" className="form-control" name="skills" placeholder="My Skills" required="" onChange={this.handleChange} value={this.state.skills} />
                                                <div className="invalid-feedback">
                                                    Please enter your skills.
                                                </div>
                                            </div>

                                            <div className="mb-3">
                                                <label htmlFor="address2">Experience <span className="text-muted">(Optional)</span></label>
                                            </div>

                                            <div className="row">
                                                <div className="col-md-5 mb-3">
                                                    <label htmlFor="zip">Company</label>
                                                    <input type="text" className="form-control" name="company" placeholder="Company" required="" onChange={this.handleChange} value={this.state.company} />
                                                    <div className="invalid-feedback">
                                                        company required.
                                                    </div>
                                                </div>
                                                <div className="col-md-5 mb-3">
                                                    <label htmlFor="zip">Role</label>
                                                    <input type="text" className="form-control" name="role" placeholder="Role" required="" onChange={this.handleChange} value={this.state.role} />
                                                    <div className="invalid-feedback">
                                                        Role required.
                                                    </div>
                                                </div>
                                            </div>
                                            <div className="col-8 mb-5">
                                                <label htmlFor="address">Role Description</label>
                                                <input type="text" className="form-control" name="roledesc" placeholder="Role Description" required="" onChange={this.handleChange} value={this.state.roledesc} />
                                                <div className="invalid-feedback">
                                                    Please enter your Role Description.
                                                </div>
                                            </div>

                                            <div className="mb-3">
                                                <label htmlFor="address2">Education <span className="text-muted">(Optional)</span></label>
                                            </div>

                                            <div className="row">
                                                <div className="col-md-5 mb-3">
                                                    <label htmlFor="zip">School</label>
                                                    <input type="text" className="form-control" name="school" placeholder="CompanSchooly" required="" onChange={this.handleChange} value={this.state.school} />
                                                    <div className="invalid-feedback">
                                                        School required.
                                                    </div>
                                                </div>
                                                <div className="col-md-5 mb-3">
                                                    <label htmlFor="zip">Program</label>
                                                    <input type="text" className="form-control" name="program" placeholder="Program" required="" onChange={this.handleChange} value={this.state.program} />
                                                    <div className="invalid-feedback">
                                                        Program required.
                                                    </div>
                                                </div>
                                            </div>
                                            <div className="col-8 mb-3">
                                                <label htmlFor="address">Program Description</label>
                                                <input type="text" className="form-control" name="progdesc" placeholder="Program Description" required="" onChange={this.handleChange} value={this.state.progdesc} />
                                                <div className="invalid-feedback">
                                                    Please enter your Program Description.
                                                </div>
                                            </div>

                                            <hr className="mb-4" />
                                            <button className="btn btn-primary btn-lg btn-block" type="submit" onClick={this.submit}>Save Profile</button>
                                        </form>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
        )
    }
}

export default Profile;