const { GraphQLServer } = require('graphql-yoga')

let links = [{
	id: 'link-0',
	url: 'www.howtographql.com',
	description: 'Fullstack tutorial for GraphQL'
}]
let count = links.length;

const resolvers = {
	Query: {
		info: () => 'This is the API of a Hackernews Clone',
		feed: () => links,
		link: (parent, args) => {
			console.log("args", args);
			return links.find((link) => link.id === args.id);
		}
	},
	Link: {
		id: (parent) => parent.id,
		description: (parent) => parent.description,
		url: (parent) => parent.url,
	},
	Mutation: {
		post: (parent, args) => {
			const link = {
				id: `link-${count++}`,
				url: args.url,
				description: args.description
			};
			links.push(link);
			return link;
		},
		updateLink: (parent, args) => {
			const i = links.findIndex((link) => link.id === args.id);
			if (i < 0) {
				return null;
			}

			const link = links[i];
			link.url = args.url || link.url;
			link.description = args.description || link.description;
			links[i] = link;
			return link
		},
		deleteLink: (parent, args) => {
			const i = links.findIndex((link) => link.id === args.id);
			if (i < 0) {
				return null;
			}

			const link = links[i];
			links.splice(i, 1);
			count = links.length;
			return link;
		}
	}
}

const server = new GraphQLServer({
	typeDefs: './src/schema.graphql',
	resolvers
})

server.start(() => console.log(`Server is running on http://localhost:4000`))
