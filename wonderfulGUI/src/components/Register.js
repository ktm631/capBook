import React, { Component } from "react";
import { Button, FormGroup, FormControl, FormLabel } from "react-bootstrap";
import "./Login.css";
import axios from 'axios';
const axiosGraphQL = axios.create({
    baseURL: 'http://localhost:7777/query',
  });

export default class Register extends Component {
  constructor(props) {
    super(props);

    this.state = {
      email: "",
      password: "",
      name: "",
      surname: ""
    };
    this.error = <span></span>
  }

  validateForm() {
    return this.state.email.length > 0 && this.state.password.length > 0 && this.state.name.length > 0 && this.state.surname.length > 0;
  }

  handleChange = event => {
    this.setState({
      [event.target.id]: event.target.value
    });
  }
  login() {
    
    axiosGraphQL.post('', {query: `
    mutation{
      createUser(input: {name: "${this.state.name}", surname: "${this.state.surname}", email: "${this.state.email}", password: "${this.state.password}", is_admin: false} ){
        name name
        surname surname
        email email
        password password
        is_admin is_admin
      }
    }

    `}).then(_ => {
        window.location.href = '/login';
    })
  }

  render() {
    return (
      <div className="register">
          <FormGroup controlId="email" bsSize="large">
            <FormLabel>Email</FormLabel><br></br>
            <FormControl
              autoFocus
              type="email"
              value={this.state.email}
              onChange={this.handleChange}
            />
          </FormGroup><br></br>
          <FormGroup controlId="name" bsSize="large">
            <FormLabel>Name</FormLabel><br></br>
            <FormControl
              value={this.state.name}
              onChange={this.handleChange}
              type="text"
            />
          </FormGroup><br></br>
          <FormGroup controlId="surname" bsSize="large">
            <FormLabel>Surname</FormLabel><br></br>
            <FormControl
              value={this.state.surname}
              onChange={this.handleChange}
              type="text"
            />
          </FormGroup><br></br>
          <FormGroup controlId="password" bsSize="large">
            <FormLabel>Password</FormLabel><br></br>
            <FormControl
              value={this.state.password}
              onChange={this.handleChange}
              type="password"
            />
          </FormGroup><br></br>
          <Button
            block
            bsSize="large"
            disabled={!this.validateForm()}
            type="submit"
            onClick={() => this.login()}
          >
            Register
          </Button><br></br>
          {this.error}
      </div>
    );
  }
}