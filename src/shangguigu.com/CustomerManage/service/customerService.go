package service

import(
	"shangguigu.com/CustomerManage/model"
)

//该CustomerService完成对customer的操作，包括增删改查
type CustomerService struct{
	//声明一个切片
	customers []model.Customer
	//声明一个字段，表示当前切片有多少个客户
	customerNum int		//该字段还可以做为新客户的Id+1
}

//编写一个方法，可以返回*CustomerService
func NewCustomerService() *CustomerService{
	//为了能够看到客户在切片中，初始化一个客户
	customerService:=&CustomerService{}
	customerService.customerNum=1
	customer:=model.NewCustomer(1,"Jack","男",20,"112","Jack@qq.com")
	customerService.customers=append(customerService.customers,customer)
	return customerService

}

//返回客户切片
func(this *CustomerService) List()[]model.Customer{
	return this.customers
}


//添加客户到customer切片
func(this *CustomerService)Add(customer model.Customer)bool{
	this.customerNum++
	customer.Id=this.customerNum
	this.customers=append(this.customers,customer)
	return true
}