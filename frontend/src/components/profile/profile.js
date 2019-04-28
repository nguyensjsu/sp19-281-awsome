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
            userid : '',
            firstname : '',
            lastname : '',
            aboutme : '',
            address : '',
            address2 : '',
            country : '',
            state : '',
            zip : '',
            company : '',
            gender : '',
            school : '',
            hometown : '',
            languages : '',
            selectedFile : [],
            profilepic : ''
        }

        this.handleChange = this.handleChange.bind(this);
        this.submit = this.submit.bind(this);
    }

    componentDidMount()
    {
        axios.get(`${USER_API}/user/`).then((response)=>{
            if(response.data){
                this.setState({
                    // firstname : response.data.users.firstname,
                    // lastname : response.data.users.lastname,
                    // aboutme: response.data.users.aboutme,
                    // address: response.data.users.address,
                    // address2: response.data.users.address2,
                    // country: response.data.users.country,
                    // state: response.data.users.state,
                    // zip: response.data.users.zip,
                    // company: response.data.users.company,
                    // gender: response.data.users.gender,
                    // school: response.data.users.school,
                    // hometown: response.data.users.hometown,
                    // languages: response.data.users.languages
                })
                //alert(response.data);
            }
        });

    }

    submit = (e) =>
    {
        e.preventDefault();

        const profiledata = {
            userid : this.state.userid,
            firstname : this.state.firstname,
            lastname : this.state.lastname,
            aboutme: this.state.aboutme,
            address: this.state.address,
            address2: this.state.address2,
            country: this.state.country,
            state: this.state.state,
            zip: this.state.zip,
            company: this.state.company,
            gender: this.state.gender,
            school: this.state.school,
            hometown: this.state.hometown,
            languages: this.state.languages,
        }


        console.log(JSON.stringify(profiledata));
        axios.put(`${USER_API}/user/`, profiledata)
            .then((response) => {
                if (response.status === 200) {
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
                                                        <p>"I throw myself down among the tall grass by the trickling
                                                            stream; and, as I lie close to the earth, a thousand unknown
                                                            plants are noticed by me: when I hear the buzz of the little
                                                            world among the stalks."</p>
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
                                                <label htmlFor="address">Address</label>
                                                <input type="text" className="form-control" name="address" placeholder="1234 Main St" required="" onChange={this.handleChange} value={this.state.address} />
                                                <div className="invalid-feedback">
                                                    Please enter your shipping address.
                                                </div>
                                            </div>

                                            <div className="mb-3">
                                                <label htmlFor="address2">Address 2 <span className="text-muted">(Optional)</span></label>
                                                <input type="text" className="form-control" name="address2" placeholder="Apartment or suite" onChange={this.handleChange} value={this.state.address2} />
                                            </div>

                                            <div className="row">
                                                <div className="col-md-5 mb-3">
                                                    <label htmlFor="country">Country</label>
                                                    <select className="custom-select d-block w-100" name="country" required="" onChange={this.handleChange} value={this.state.country} >
                                                        <option value="">Choose...</option>
                                                        <option value="UnitedState">United States</option>
                                                    </select>
                                                    <div className="invalid-feedback">
                                                        Please select a valid country.
                                                    </div>
                                                </div>
                                                <div className="col-md-4 mb-3">
                                                    <label htmlFor="state">State</label>
                                                    <select className="custom-select d-block w-100" name="state" required="" onChange={this.handleChange} value={this.state.state} >
                                                        <option value="">Choose...</option>
                                                        <option value="California">California</option>
                                                    </select>
                                                    <div className="invalid-feedback">
                                                        Please provide a valid state.
                                                    </div>
                                                </div>
                                                <div className="col-md-3 mb-3">
                                                    <label htmlFor="zip">Zip</label>
                                                    <input type="text" className="form-control" name="zip" placeholder="" required="" onChange={this.handleChange} value={this.state.zip} />
                                                    <div className="invalid-feedback">
                                                        Zip code required.
                                                    </div>
                                                </div>
                                            </div>
                                            <div className="row">
                                                <div className="col-md-5 mb-5">
                                                    <label htmlFor="zip">Company</label>
                                                    <input type="text" className="form-control" name="company" placeholder="Company" required="" onChange={this.handleChange} value={this.state.company} />
                                                    <div className="invalid-feedback">
                                                        company required.
                                                    </div>
                                                </div>
                                                <div className="col-md-4 mb-3">
                                                    <label htmlFor="state">Gender</label>
                                                    <select className="custom-select d-block w-100" name="gender" required="" onChange={this.handleChange} value={this.state.gender} >
                                                        <option value="">Choose...</option>
                                                        <option value="Male">Male</option>
                                                        <option value="Female">Female</option>
                                                        <option value="No">I Don't want to disclose my Identity</option>
                                                    </select>
                                                    <div className="invalid-feedback">
                                                        Please provide a valid state.
                                                    </div>
                                                </div>
                                            </div>
                                            <div className="row">
                                                <div className="col-md-5 mb-5">
                                                    <label htmlFor="zip">School</label>
                                                    <input type="text" className="form-control" name="school" placeholder="School" required="" onChange={this.handleChange} value={this.state.school}/>
                                                    <div className="invalid-feedback">
                                                        School required.
                                                    </div>
                                                </div>

                                                <div className="col-md-5 mb-5">
                                                    <label htmlFor="zip">HomeTown</label>
                                                    <input type="text" className="form-control" name="hometown" placeholder="Hometown" required="" onChange={this.handleChange} value={this.state.hometown} />
                                                    <div className="invalid-feedback">
                                                        Hometown required.
                                                    </div>
                                                </div>
                                            </div>
                                            <div className="col-md-8 mb-5">
                                                <label htmlFor="zip">Languages</label>
                                                <input type="text" className="form-control" name="languages" placeholder="Languages" required="" onChange={this.handleChange} value={this.state.languages} />
                                                <div className="invalid-feedback">
                                                    Languages required.
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