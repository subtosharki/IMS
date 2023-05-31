export interface User {
	id: string;
	userId: string;
	apikey: string;
	email: string;
	firstName: string;
	lastName: string;
	password: string;
	role: 'sales' | 'admin';
}

export interface CreateUserBody {
	email: string;
	firstName: string;
	lastName: string;
	password: string;
	admin: boolean;
}

export type OrderStatus = 'Voided' | 'Fulfilled' | 'In Progress';

export interface Clone {
	id: string;
	cloneId: string;
	name: string;
	quantity: number;
	date: string;
}

export interface OrderClone extends Clone {
	cloneId: string;
	selectedQuantity: number;
}

export interface Order {
	'cloneId': string;
	'use': string;
	'customerName': string;
	'datePlaced': string;
	'dateRequired': string;
	'orderNumber': string;
	'clones': Clone[];
	'status': string;
	'notes': string;
	'placedBy': string;
}

export interface Customer {
	id: string;
	customerId: string;
	name: string;
	email: string;
	subscriptionId: string;
	phone: string;
	city: string;
	state: string;
	zip: string;
	notes: string;
}

export interface CreateCustomerBody {
	name: string;
	email: string;
	subscriptionId: string;
	phone: string;
	city: string;
	state: string;
	zip: string;
}
