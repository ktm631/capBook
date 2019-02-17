import React from 'react';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import { Component } from "react";
import axios from 'axios';
const axiosGraphQL = axios.create({
    baseURL: 'http://localhost:7777/query',
  });


class BookElement extends Component {
  state = {
    json: {},
    loading: 0,
};
    componentDidMount() {
        this.onFetch();
      }
    onFetch() {
        axiosGraphQL
        .post('', { query: `
        query{
          book(book_id: ${this.props.match.params.id}) {
            title
            author_id
            location_id
            owner_id
            edition
            isbn
            is_free
            publisher_id
            description_url
          }
        }
        ` })
        .then(result =>  {
            this.setState({ json: result, loading: 0 })
            return result;
          })
          .then(result => {
            axiosGraphQL.post('', {
              query: `
              query{
                author(author_id: ${result.data.data.book.author_id}){
                  author_id
                  name
                  surname
                }
              }
              `
            }).then(e => {

                this.author = `${e.data.data.author.name} ${e.data.data.author.surname}`
              return true;
            })
            .then(_ => this.setState({ loading: this.state.loading + 1 }));
            
            return result;
          })
          .then(result => {
            axiosGraphQL.post('', {
              query: `
              query{
                location(location_id: ${result.data.data.book.location_id}){
                  location_id
                  building
                  room
                }
              }
              `
            }).then(e => {

                this.location = `${e.data.data.location.building} ${e.data.data.location.room}`
              return true;
            })
            .then(_ => this.setState({ loading: this.state.loading + 1 }));
            
            return result;
          })
          .then(result => {
            axiosGraphQL.post('', {
              query: `
              query{
                user(user_id: ${result.data.data.book.owner_id}){
                user_id
                  name
                  surname
                }
              }
              `
            }).then(e => {

                this.owner = `${e.data.data.user.name} ${e.data.data.user.surname}`
              return true;
            })
            .then(_ => this.setState({ loading: this.state.loading + 1 }));
            
            return result;
          })
          .then(result => {
            axiosGraphQL.post('', {
              query: `
              query{
                publisher(publisher_id: ${result.data.data.book.publisher_id}){
                  publisher_id
                  name
                }
                }
              `
            }).then(e => {

                this.publisher = `${e.data.data.publisher.name}`
              return true;
            })
            .then(_ => this.setState({ loading: this.state.loading + 1 }));
            
            return result;
          });
    }
    handleClick() {
console.log('click');

    }
    render() {
    const { loading, json } = this.state;
    if (loading < 4) {
        return <span>Loading...</span>
    }
     return (
    <Card className="register left order">
      <CardContent>
        <Typography variant="h5">
        Tytuł: {json.data.data.book.title}
        </Typography>
        <Typography variant="h5">
        Autor: {this.author}
        </Typography>
        <Typography variant="h5">
        Właściciel: {this.owner}
        </Typography>
        <Typography variant="h5">
        Edycja: {json.data.data.book.edition}
        </Typography>
        <Typography variant="h5">
        ISBN: {json.data.data.book.isbn}
        </Typography>
        <Typography variant="h5">
        Lokacja: {this.location}
        </Typography>
        <Typography variant="h5">
        Wydawca: {this.publisher}
        </Typography>
        <Typography variant="h5">
        Link: <a href={json.data.data.book.description_url}>...</a>
        </Typography>
      <Button
            bsSize="large"
            onClick={this.handleClick} disabled={!json.data.data.book.is_free}>Wypożycz</Button>
      </CardContent>
    </Card>
  );}
}

export default BookElement;