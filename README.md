<h1 align="center">Apache Fineract Golang SDK</h3>

[![Godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/wondenge/go-fineract?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/wondenge/coop-go)](https://goreportcard.com/report/wondenge/go-fineract)
[![License](https://img.shields.io/badge/license-Apache-blue.svg)]

# Overview

Apache Fineract (\’fīn-,ә-,rakt\) is open source software for financial services.

Fineract provides a reliable, robust, and affordable solution for entrepreneurs, financial institutions, and service providers to offer financial services to the world’s 2 billion underbanked and unbanked. Fineract is aimed at innovative mobile and cloud-based solutions, and enables digital transaction accounts for all.

Fineract 1.x compares well to other core banking systems and draws from requirements in credit unions, microfinance institutions, and small non-banking financial institutions. Features include flexible product configuration, KYC documentation support, business rule sets, payment transactions, and portfolio management. It includes an open API that dates to 2011 and is deployed in relatively high transaction volume environments.

# Installation

```
go get github.com/wondenge/go-fineract
```

# Usage

## Periodic Accrual Accounting

Periodic Accrual is to accrue the loan income till the specific date or till batch job scheduled time.

## Accounting Closure

An accounting closure indicates that no more journal entries may be logged (or reversed) in the system, either manually or via the portfolio with an entry date prior to the defined closure date

Field Descriptions closingDate The date for which the accounting closure is defined officeId The identifer of the branch for which accounting has been closed comments Description associated with an Accounting closure

## Mapping Financial Activities to Accounts

Organization Level Financial Activities like Asset and Liability Transfer can be mapped to GL Account. Integrated accounting takes these accounts into consideration when an Account transfer is made between a savings to loan/savings account and vice-versa

Field Descriptions financialActivityId The identifier of the Financial Activity glAccountId The identifier of a GL Account ( Ledger Account) which shall be used as the default account for the selected Financial Activity

## General Ledger Account

Ledger accounts represent an Individual account within an Organizations Chart Of Accounts(COA) and are assigned a name and unique number by which they can be identified. All transactions relating to a company's assets, liabilities, owners' equity, revenue and expenses are recorded against these accounts

## Journal Entries

A journal entry refers to the logging of transactions against general ledger accounts. A journal entry may consist of several line items, each of which is either a "debit" or a "credit". The total amount of the debits must equal the total amount of the credits or the journal entry is said to be "unbalanced"

A journal entry directly changes the account balances on the general ledger

## Provisioning Entries

This defines the Provisioning Entries for all active loan products

Field Descriptions date Date on which day provisioning entries should be created createjournalentries Boolean variable whether to add journal entries for generated provisioning entries

## Accounting Rules

It is typical scenario in MFI's that non accountants pass journal entries on a regular basis. For Ex: A branch office might deposit their entire cash at hand to their Bank account at the end of a working day. The branch office users might not understand enough of accounting to figure out which account needs to get credited and which account needs to be debited to represent this transaction.

Enter accounting rules, an abstraction on top of manual Journal entires for enabling simpler data entry. An accounting rule can define any of the following abstractions

A Simple journal entry where both the credit and debit account have been preselected A Simple journal entry where either credit or debit accounts have been limited to a pre-selected list of accounts (Ex: Debit account should be one of "Bank of America" of "JP Morgan" and credit account should be "Cash") A Compound journal entry where multiple debits and / or multiple credits may be made amongst a set of preselected list of accounts (Ex: Credit account should be either "Bank Of America" or "Cash" and debit account can be "Employee Salary" and/or "Miscellenous Expenses") An accounting rule can also be optionally associated with a branch, so that only a particular Branch's users have access to the rule

## Batch API

The Apache Fineract Batch API enables a consumer to access significant amounts of data in a single call or to make changes to several objects at once. Batching allows a consumer to pass instructions for several operations in a single HTTP request. A consumer can also specify dependencies between related operations. Once all operations have been completed, a consolidated response will be passed back and the HTTP connection will be closed.

The Batch API takes in an array of logical HTTP requests represented as JSON arrays - each request has a requestId (the id of a request used to specify the sequence and as a dependency between requests), a method (corresponding to HTTP method GET/PUT/POST/DELETE etc.), a relativeUrl (the portion of the URL after https://example.org/api/v2/), optional headers array (corresponding to HTTP headers), optional reference parameter if a request is dependent on another request and an optional body (for POST and PUT requests). The Batch API returns an array of logical HTTP responses represented as JSON arrays - each response has a requestId, a status code, an optional headers array and an optional body (which is a JSON encoded string).

Batch API uses Json Path to handle dependent parameters. For example, if request '2' is referencing request '1' and in the "body" or in "relativeUrl" of request '2', there is a dependent parameter (which will look like "\$.parameter_name"), then Batch API will internally substitute this dependent parameter from the response body of request '1'.

Batch API is able to handle deeply nested dependent requests as well nested parameters. As shown in the example, requests are dependent on each other as, 1<--2<--6, i.e a nested dependency, where request '6' is not directly dependent on request '1' but still it is one of the nested child of request '1'. In the same way Batch API could handle a deeply nested dependent value, such as {..[..{..,$.parameter_name,..}..]}.

## Audits

Every non-read Mifos API request is audited. A fully processed request can not be changed or deleted. See maker checker api for situations where an audit is not fully processed.

Permissions: To search and look at audit entries a user needs to be attached to a role that has one of the ALL_FUNCTIONS, ALL_FUNCTIONS_READ or READ_AUDIT permissions.

Data Scope: A user can only see audits that are within their data scope. However, 'head office' users can see all audits including those that aren't office/branch related e.g. Loan Product changes.")

## Account number format

Account number preferences are used to describe custom formats for account numbers associated with Customer, Loan and Savings accounts.

## Cache

The following settings are possible for cache:

No Caching: caching turned off Single node: caching on for single instance deployments of platorm (works for multiple tenants but only one tomcat) By default caching is set to No Caching. Switching between caches results in the cache been clear e.g. from Single node to No cache and back again would clear down the single node cache.

## Code Values

Code and code values: Codes represent a specific category of data, their code values are a specific instance of that category.

Codes are mostly system defined which means the code itself comes out of the box and cannot be modified however its code values can be. e.g. 'Customer Identifier', it defaults to a code value of 'Passport' but could be 'Drivers License, National Id' etc

## Codes

Code and code values: Codes represent a specific category of data, their code values are a specific instance of that category.

Codes are mostly system defined which means the code itself comes out of the box and cannot be modified however its code values can be. e.g. 'Customer Identifier', it defaults to a code value of 'Passport' but could be 'Drivers License, National Id' etc

## External Services

External Services Configuration related to set of supported configurations for third party services like Amazon S3 and SMTP:

S3 (Amazon S3): s3_access_key - s3_bucket_name - s3_secret_key -

SMTP (Email Service): username - password - host - port - useTLS -

## Global Configuration

Global configuration related to set of supported enable/disable configurations:

maker-checker - defaults to false - if true turns on maker-checker functionality reschedule-future-repayments - defaults to false - if true reschedules repayemnts which falls on a non-working day to configured repayment rescheduling rule allow-transactions-on-non_workingday - defaults to false - if true allows transactions on non-working days reschedule-repayments-on-holidays - defaults to false - if true reschedules repayemnts which falls on a non-working day to defined reschedule date allow-transactions-on-holiday - defaults to false - if true allows transactions on holidays savings-interest-posting-current-period-end - Set it at the database level before any savings interest is posted. When set as false(default), interest will be posted on the first date of next period. If set as true, interest will be posted on last date of current period. There is no difference in the interest amount posted. financial-year-beginning-month - Set it at the database level before any savings interest is posted. Allowed values 1 - 12 (January - December). Interest posting periods are evaluated based on this configuration. meetings-mandatory-for-jlg-loans - if set to true, enforces all JLG loans to follow a meeting schedule belonging to either the parent group or Center.

## Data Tables

The datatables API allows you to plug-in your own tables (MySql) that have a relationship to a Apache Fineract core table. For example, you might want to add some extra client fields and record information about each of the clients' family members. Via the API you can create, read, update and delete entries for each 'plugged-in' table. The API checks for permission and for 'data scoping' (only data within the users' office hierarchy can be managed by the user).

The Apache Fineract Reference App uses a JQuery plug-in called stretchydatatables (which in turn uses this datatables resource) to provide a pretty flexible CRUD (Create, Read, Update, Delete) User Interface.

## Entity Data Table

This defines Entity-Datatable Check.

## Reports

Non-core reports can be added, updated and deleted.

## Documents

Multiple Documents (a combination of a name, description and a file) may be attached to different Entities like Clients, Groups, Staff, Loans, Savings and Client Identifiers in the system

Note: The currently allowed Entities are

Clients: URL Pattern as clients Staff: URL Pattern as staff Loans: URL Pattern as loans Savings: URL Pattern as savings Client Identifiers: URL Pattern as client_identifiers Groups: URL Pattern as groups

## Hooks

Hooks are a mechanism to trigger custom code on the occurence of events.

## MIFOSX-BATCH JOBS

Batch jobs (also known as cron jobs on Unix-based systems) are a series of back-end jobs executed on a computer at a particular time defined in job's cron expression.

At any point, you can view the list of batch jobs scheduled to run along with other details specific to each job. Manually you can execute the jobs at any point of time.

The scheduler status can be either "Active" or "Standby". If the scheduler status is Active, it indicates that all batch jobs are running/ will run as per the specified schedule.If the scheduler status is Standby, it will ensure all scheduled batch runs are suspended.

## Report Mailing Jobs

This resource allows you to create a scheduled job that runs a report and sents it by email to specified email addresses.

The scheduled job can be configured to run once or on a regular basis (once a day, twice a week, etc)

## Authentication HTTP Basic

An API capability that allows client applications to verify authentication details using HTTP Basic Authentication.

## Holidays

Some MFI's span large regions where different branch offices might observe different holidays. They need the ability to define holidays for specific branch offices and be able to set the repayment rule to follow during those holidays.

The reschedule of repayments to repaymentsRescheduledTo date during defined holidays is turned on/off by enabling/disabling reschedule-repayments-on-holidays in Global configurations.

Allow Repayment transactions on a defined holidays is turned on/off by enabling/disabling allow-transactions-on-holiday in Global configurations.

## Currency

Application related configuration around viewing/updating the currencies permitted for use within the MFI.

## Offices

Offices are used to model an MFIs structure. A hierarchical representation of offices is supported. There will always be at least one office (which represents the MFI or an MFIs head office). All subsequent offices added must have a parent office.

## Provisioning Criteria

This defines the Provisioning Criteria

## Staff

Allows you to model staff members. At present the key role of significance is whether this staff member is a loan officer or not.

## Teller Cash Management

Teller cash management which will allow an organization to manage their cash transactions at branches or head office more effectively.

## Working days

The days of the week that are workdays.

Rescheduling of repayments when it falls on a non-working is turned on /off by enable/disable reschedule-future-repayments parameter in Global configurations

Allow transactions on non-working days is configurable by enabling/disbaling the allow-transactions-on-non_workingday parameter in Global configurations.

## Account Transfers

Ability to be able to transfer monetary funds from one account to another.

Note: At present only savings account to savings account transfers are supported.

## Standing Instructions

Standing instructions (or standing orders) refer to instructions a bank account holder ("the payer") gives to his or her bank to pay a set amount at regular intervals to another's ("the payee's") account.

Note: At present only savings account to savings account and savings account to Loan account transfers are permitted.

## Standing Instructions History

The list capability of history can support pagination and sorting.

## Share Account

Share accounts are instances of a praticular share product created for an individual. An application process around the creation of accounts is also supported.

## Entity Field Configuration

Entity Field configuration API is a generic and extensible wherein various entities and subentities can be related. Also it gives the user an ability to enable/disable fields, add regular expression for validation

## Charges

Its typical for MFIs to add extra costs for their financial products. These are typically Fees or Penalties.

A Charge on fineract platform is what we use to model both Fees and Penalties.

At present we support defining charges for use with Client accounts and both loan and saving products.

## Clients Address

Address module is an optional module and can be configured into the system by using GlobalConfiguration setting: enable-address. In order to activate Address module, we need to enable the configuration, enable-address by setting its value to true.

## Client Charges

It is typical for MFI's to directly associate charges with an implicit Client account. These can be either fees or penalties

Client Charges are client specific instances of Charges. Refer Charges for documentation of the various properties of a charge, Only additional properties ( specific to the context of a Charge being associated with a Client account) are described here

## Client Identifier

Client Identifiers refer to documents that are used to uniquely identify a customer Ex: Drivers License, Passport, Ration card etc

## Client Transaction

Client Transactions refer to transactions made directly againt a Client's internal account. Currently, these transactions are only created as a result of charge payments/waivers. You are allowed to undo a transaction, however you cannot explicitly create one.

## Client

Clients are people and businesses that have applied (or may apply) to an MFI for loans.

Clients can be created in Pending or straight into Active state.

## Loan Collateral

In lending agreements, collateral is a borrower's pledge of specific property to a lender, to secure repayment of a loan. The collateral serves as protection for a lender against a borrower's default - that is, any borrower failing to pay the principal and interest under the terms of a loan obligation. If a borrower does default on a loan (due to insolvency or other event), that borrower forfeits (gives up) the property pledged as collateral - and the lender then becomes the owner of the collateral

## Floating Rates

It lets you create, list, retrieve and upload the floating rates

## Centers

Centers along with Groups are used to provided a distinctive banking distribution channel used in microfinance. Its common in areas such as Southern Asia to use Centers and Group as administrative units in grameen style lending. Typically groups will contain one to five people and centers themselves will be made of anywhere between 2-10 groups.

## Groups

Groups are used to provide a distinctive banking distribution channel used in microfinances throughout the world. The Group is an administrative unit. It can contain as few as 5 people or as many as 40 depending on how its used.

Different styles of group lending - Joint-Liability Group, Grameen Model (Center-Group), Self-Help Groups, Village/Communal Banks)

## Interest Rate Slab (A.K.A interest bands)

The slabs a.k.a interest bands are associated with Interest Rate Chart. These bands allow to define different interest rates for different deposit term periods.

## Interest Rate Chart

This defines an interest rate scheme that can be associated to a term deposit product. This will have a slab (band or range) of deposit periods and the associated interest rates applicable along with incentives for each band.

## Loan Charges

Its typical for MFIs to add extra costs for their loan products. They can be either Fees or Penalties.

Loan Charges are instances of Charges and represent either fees and penalties for loan products. Refer Charges for documentation of the various properties of a charge, Only additional properties ( specific to the context of a Charge being associated with a Loan) are described here

## Loan Rescheduling

Loan Term Variations provides the ability to change due dates, amounts and number of instalments before loan approval.

## Loan Transactions

Capabilities include loan repayment's, interest waivers and the ability to 'adjust' an existing transaction. An 'adjustment' of a transaction is really a 'reversal' of existing transaction followed by creation of a new transaction with the provided details.

## Loans

The API concept of loans models the loan application process and the loan contract/monitoring process.

Field Descriptions accountNo The account no. associated with this loan. Is auto generated if not provided at loan application creation time. externalId A place to put an external reference for this loan e.g. The ID another system uses. If provided, it must be unique. fundId Optional: For associating a loan with a given fund. loanOfficerId Optional: For associating a loan with a given staff member who is a loan officer. loanPurposeId Optional: For marking a loan with a given loan purpose option. Loan purposes are configurable and can be setup by system admin through code/code values screens. principal The loan amount to be disbursed to through loan. loanTermFrequency The length of loan term Used like: loanTermFrequency loanTermFrequencyType e.g. 12 Months loanTermFrequencyType The loan term period to use. Used like: loanTermFrequency loanTermFrequencyType e.g. 12 Months Example Values: 0=Days, 1=Weeks, 2=Months, 3=Years numberOfRepayments Number of installments to repay. Used like: numberOfRepayments Every repaymentEvery repaymentFrequencyType e.g. 10 (repayments) Every 12 Weeks repaymentEvery Used like: numberOfRepayments Every repaymentEvery repaymentFrequencyType e.g. 10 (repayments) Every 12 Weeks repaymentFrequencyType Used like: numberOfRepayments Every repaymentEvery repaymentFrequencyType e.g. 10 (repayments) Every 12 Weeks Example Values: 0=Days, 1=Weeks, 2=Months interestRatePerPeriod Interest Rate. Used like: interestRatePerPeriod % interestRateFrequencyType - interestType e.g. 12.0000% Per year - Declining Balance interestRateFrequencyType Used like: interestRatePerPeriod% interestRateFrequencyType - interestType e.g. 12.0000% Per year - Declining Balance Example Values: 2=Per month, 3=Per year graceOnPrincipalPayment Optional: Integer - represents the number of repayment periods that grace should apply to the principal component of a repayment period. graceOnInterestPayment Optional: Integer - represents the number of repayment periods that grace should apply to the interest component of a repayment period. Interest is still calculated but offset to later repayment periods. graceOnInterestCharged Optional: Integer - represents the number of repayment periods that should be interest-free. graceOnArrearsAgeing Optional: Integer - Used in Arrears calculation to only take into account loans that are more than graceOnArrearsAgeing days overdue. interestChargedFromDate Optional: Date - The date from with interest is to start being charged. expectedDisbursementDate The proposed disbursement date of the loan so a proposed repayment schedule can be provided. submittedOnDate The date the loan application was submitted by applicant. linkAccountId The Savings Account id for linking with loan account for payments. amortizationType Example Values: 0=Equal principle payments, 1=Equal installments interestType Used like: interestRatePerPeriod% interestRateFrequencyType - interestType e.g. 12.0000% Per year - Declining Balance Example Values: 0=Declining Balance, 1=Flat interestCalculationPeriodType Example Values: 0=Daily, 1=Same as repayment period allowPartialPeriodInterestCalcualtion This value will be supported along with interestCalculationPeriodType as Same as repayment period to calculate interest for partial periods. Example: Interest charged from is 5th of April , Principal is 10000 and interest is 1% per month then the interest will be (10000 _ 1%)_ (25/30) , it calculates for the month first then calculates exact periods between start date and end date(can be a decimal) inArrearsTolerance The amount that can be 'waived' at end of all loan payments because it is too small to worry about. This is also the tolerance amount assessed when determining if a loan is in arrears. transactionProcessingStrategyId An enumeration that indicates the type of transaction processing strategy to be used. This relates to functionality that is also known as Payment Application Logic. A number of out of the box approaches exist, some are custom to specific MFIs, some are more general and indicate the order in which payments are processed.

Refer to the Payment Application Logic / Transaction Processing Strategy section in the appendix for more detailed overview of each available payment application logic provided out of the box.

List of current approaches: 1 = Mifos style (Similar to Old Mifos) 2 = Heavensfamily (Custom MFI approach) 3 = Creocore (Custom MFI approach) 4 = RBI (India) 5 = Principal Interest Penalties Fees Order 6 = Interest Principal Penalties Fees Order 7 = Early Payment Strategy loanType To represent different type of loans. At present there are three type of loans are supported. Available loan types: individual: Loan given to individual member group: Loan given to group as a whole jlg: Joint liability group loan given to members in a group on individual basis. JLG loan can be given to one or more members in a group. recalculationRestFrequencyDate Specifies rest frequency start date for interest recalculation. This date must be before or equal to disbursement date recalculationCompoundingFrequencyDate Specifies compounding frequency start date for interest recalculation. This date must be equal to disbursement date

## Loan Products

A Loan product is a template that is used when creating a loan. Much of the template definition can be overridden during loan creation.

## Notes

Notes API allows to enter notes for supported resources.

## Payment Type

This defines the payment type

## Fixed Deposit Account

Fixed Deposit accounts are instances of a praticular fixed deposit product created. An application process around the creation of accounts is also supported.

## Fixed Deposit Product

This is one of the advanced term deposit product offered by MFI's. The Fixed Deposit Products (aka FD) product offerings are modeled using this API.

The FD products are deposit accounts which are held for a fixed term – like 1 year, 2 years etc.

When creating fixed deposit accounts, the details from the fixed deposit product are used to auto fill details of the fixed deposit account application process.

## Recurring Deposit Account Transactions

Transactions possible on a recurring deposit account.

## Recurring Deposit Account

Recurring Deposit accounts are instances of a praticular recurring deposit product created. An application process around the creation of accounts is also supported.

## Recurring Deposit Product

Recurring Deposits are a special kind of Term Deposits offered by MFI's. The Recurring Deposit Products (aka RD) product offerings are modeled using this API.

Recurring Deposits help people with regular incomes to deposit a fixed amount every month (specified recurring frequency) into their Recurring Deposit account.

When creating recurring deposit accounts, the details from the recurring deposit product are used to auto fill details of the recurring deposit account application process.

## Savings Charges

Its typical for MFIs to add maintenance and operating charges. They can be either Fees or Penalties.

Savings Charges are instances of Charges and represent either fees and penalties for savings products. Refer Charges for documentation of the various properties of a charge, Only additional properties ( specific to the context of a Charge being associated with a Savings account) are described here

## Savings Account

Savings accounts are instances of a particular savings product created for an individual or group. An application process around the creation of accounts is also supported.

## Savings Product

An MFIs savings product offerings are modeled using this API. When creating savings accounts, the details from the savings product are used to auto fill details of the savings account application process.

## Search API

Search API allows to search scoped resources clients, loans and groups on specified fields.

## Self Loan Products

A Loan product is a template that is used when creating a loan. Much of the template definition can be overridden during loan creation.

Field Descriptions name Name associated with loan product on system. shortName Short name associated with a loan product. An abbreviated version of the name, used in reports or menus where space is limited, such as Collection Sheets. description For providing helpful description of product offering. fundId For associating a loan product with a given fund by default. includeInBorrowerCycle It is a flag, Used to denote whether the loans should include in loan cycle counter or not. useBorrowerCycle It is a flag, Used to denote whether the loans should depend on borrower loan cycle counter or not. currencyCode A three letter ISO code of currency. digitsAfterDecimal Override the currency default value for digitsAfterDecimal. inMultiplesOf Override the default value for rounding currency to multiples of value provided. installmentAmountInMultiplesOf Override the default value for rounding instalment amount to multiples of value provided. principal The loan amount to be disbursed to through loan. numberOfRepayments Number of installments to repay. Used like: numberOfRepayments Every repaymentEvery repaymentFrequencyType e.g. 10 (repayments) Every 12 Weeks repaymentEvery Used like: numberOfRepayments Every repaymentEvery repaymentFrequencyType e.g. 10 (repayments) Every 12 Weeks repaymentFrequencyType Used like: numberOfRepayments Every repaymentEvery repaymentFrequencyType e.g. 10 (repayments) Every 12 Weeks Example Values: 0=Days, 1=Weeks, 2=Months interestRatePerPeriod Interest Rate. Used like: interestRatePerPeriod % interestRateFrequencyType - interestType e.g. 12.0000% Per year - Declining Balance interestRateFrequencyType Used like: interestRatePerPeriod% interestRateFrequencyType - interestType e.g. 12.0000% Per year - Declining Balance Example Values: 2=Per month, 3=Per year amortizationType Example Values: 0=Equal principle payments, 1=Equal installments interestType Used like: interestRatePerPeriod% interestRateFrequencyType - interestType e.g. 12.0000% Per year - Declining Balance Example Values: 0=Declining Balance, 1=Flat interestCalculationPeriodType Example Values: 0=Daily, 1=Same as repayment period allowPartialPeriodInterestCalcualtion This value will be supported along with interestCalculationPeriodType as Same as repayment period to calculate interest for partial periods. Example: Interest charged from is 5th of April , Principal is 10000 and interest is 1% per month then the interest will be (10000 _ 1%)_ (25/30) , it calculates for the month first then calculates exact periods between start date and end date(can be a decimal) inArrearsTolerance The amount that can be 'waived' at end of all loan payments because it is too small to worry about. This is also the tolerance amount assessed when determining if a loan is in arrears. principalVariationsForBorrowerCycle,interestRateVariationsForBorrowerCycle, numberOfRepaymentVariationsForBorrowerCycle Variations for loan, based on borrower cycle number minimumDaysBetweenDisbursalAndFirstRepayment The minimum number of days allowed between a Loan disbursal and its first repayment. principalThresholdForLastInstalment Field represents percentage of current instalment principal amount for comparing against principal outstanding to add another repayment instalment. If the outstanding principal amount is less then calculated amount, remaining outstanding amount will be added to current instalment. Default value for multi disburse loan is 50% and non-multi disburse loan is 0% canDefineInstallmentAmount if provided as true, then fixed instalment amount can be provided from loan account. transactionProcessingStrategyId An enumeration that indicates the type of transaction processing strategy to be used. This relates to functionality that is also known as Payment Application Logic. A number of out of the box approaches exist, some are custom to specific MFIs, some are more general and indicate the order in which payments are processed.

Refer to the Payment Application Logic / Transaction Processing Strategy section in the appendix for more detailed overview of each available payment application logic provided out of the box.

List of current approaches: 1 = Mifos style (Similar to Old Mifos) 2 = Heavensfamily (Custom MFI approach) 3 = Creocore (Custom MFI approach) 4 = RBI (India) 5 = Principal Interest Penalties Fees Order 6 = Interest Principal Penalties Fees Order 7 = Early Payment Strategy graceOnPrincipalPayment Optional: Integer - represents the number of repayment periods that grace should apply to the principal component of a repayment period. graceOnInterestPayment Optional: Integer - represents the number of repayment periods that grace should apply to the interest component of a repayment period. Interest is still calculated but offset to later repayment periods. graceOnInterestCharged Optional: Integer - represents the number of repayment periods that should be interest-free. graceOnArrearsAgeing Optional: Integer - Used in Arrears calculation to only take into account loans that are more than graceOnArrearsAgeing days overdue. overdueDaysForNPA Optional: Integer - represents the maximum number of days a Loan may be overdue before being classified as a NPA (non performing asset) accountMovesOutOfNPAOnlyOnArrearsCompletion Optional: Boolean - if provided as true, Loan Account moves out of NPA state only when all arrears are cleared accountingRule Specifies if accounting is enabled for the particular product and the type of the accounting rule to be used Example Values:1=NONE, 2=CASH_BASED, 3=ACCRUAL_PERIODIC, 4=ACCRUAL_UPFRONT isInterestRecalculationEnabled It is a flag, Used to denote whether interest recalculation is enabled or disabled for the particular product daysInYearType Specifies the number of days in a year. Example Values:1=ACTUAL(Actual number of days in year), 360=360 DAYS, 364=364 DAYS(52 WEEKS), 365=365 DAYS daysInMonthType Specifies the number of days in a month. Example Values:1=ACTUAL(Actual number of days in month), 30=30 DAYS interestRecalculationCompoundingMethod Specifies which amount portion should be added to principal for interest recalculation. Example Values:0=NONE(Only on principal), 1=INTEREST(Principal+Interest), 2=FEE(Principal+Fee), 3=FEE And INTEREST (Principal+Fee+Interest) rescheduleStrategyMethod Specifies what action should perform on loan repayment schedule for advance payments. Example Values:1=Reschedule next repayments, 2=Reduce number of installments, 3=Reduce EMI amount recalculationCompoundingFrequencyType Specifies effective date from which the compounding of interest or fee amounts will be considered in recalculation on late payment. Example Values:1=Same as repayment period, 2=Daily, 3=Weekly, 4=Monthly recalculationCompoundingFrequencyInterval Specifies compounding frequency interval for interest recalculation. recalculationCompoundingFrequencyDate Specifies compounding frequency start date for interest recalculation. recalculationRestFrequencyType Specifies effective date from which the late or advanced payment amounts will be considered in recalculation. Example Values:1=Same as repayment period, 2=Daily, 3=Weekly, 4=Monthly recalculationRestFrequencyInterval Specifies rest frequency interval for interest recalculation. recalculationRestFrequencyDate Specifies rest frequency start date for interest recalculation. preClosureInterestCalculationStrategy Specifies applicable days for interest calculation on pre closure of a loan. Example Values:1=Calculate till pre closure date, 2=Calculate till rest frequency date isArrearsBasedOnOriginalSchedule If Specified as true, arrears will be identified based on original schedule. allowAttributeOverrides Specifies if select attributes may be overridden for individual loan accounts.

## Self Run Report

This resource allows you to run and receive output from pre-defined Apache Fineract reports.

The default output is a JSON formatted "Generic Resultset". The Generic Resultset contains Column Heading as well as Data information. However, you can export to CSV format by simply adding "&exportCSV=true" to the end of your URL.

If Pentaho reports have been pre-defined, they can also be run through this resource. Pentaho reports can return HTML, PDF or CSV formats.

The Apache Fineract reference application uses a JQuery plugin called stretchyreporting which, itself, uses this reports resource to provide a pretty flexible reporting User Interface (UI).

ARGUMENTS R*'parameter names' ... optional, No defaults The number and names of the parameters depend on the specific report and how it has been configured. R_officeId is an example parameter name.Note: the prefix R* stands for ReportinggenericResultSetoptional, defaults to true If 'true' an optimised JSON format is returned suitable for tabular display of data. If 'false' a simple JSON format is returned. parameterType optional, The only valid value is 'true'. If any other value is provided the argument will be ignored Determines whether the request looks in the list of reports or the list of parameters for its data. Doesn't apply to Pentaho reports.exportCSV optional, The only valid value is 'true'. If any other value is provided the argument will be ignored Output will be delivered as a CSV file instead of JSON. Doesn't apply to Pentaho reports.output-type optional, Defaults to HTML. Valid Values are HTML, XLS, XSLX, CSV and PDF for html, Excel, Excel 2007+, CSV and PDF formats respectively.Only applies to Pentaho reports.locale optional Any valid locale Ex: en_US, en_IN, fr_FR etcOnly applies to Pentaho reports.

## Self Authentication

Authenticates the credentials provided and returns the set roles and permissions allowed

## Tax Components

This defines the Tax Components

## Tax Group

This defines the Tax Group

## SPM API - LookUp Table

The Apache Fineract SPM API provides the ability to create custom surveys to collect social performance measurentment data or any additional questionnaire a financial institute want to collect.

## User Generated Documents

User Generated Documents(alternatively, Templates) are used for end-user features such as custom user defined document generation (AKA UGD). They are based on {{ moustache }} templates. Think of them as a sort of built-in "mail merge" functionality.

User Generated Documents (and other types of templates) can aggregate data from several Apache Fineract back-end API calls via mappers. Mappers can even access non-Apache Fineract REST services from other servers. UGDs can render such data in tables, show images, etc. TBD: Please have a look at some of the Example UGDs included in Apache Fineract (or the Wiki page, for now.).

UGDs can be assigned to an entity like client or loan and be of a type like Document or SMS. The entity and type of a UGD is only there for the convenience of user agents (UIs), in order to know where to show UGDs for the user (i.e. which tab). The Template Engine back-end runner does not actually need this metadata.

## Password preferences

This API enables management of password policy for user administration.

There is no Apache Fineract functionality for creating a validation policy. The validation policies come pre-installed.

Validation policies may be updated

## Permissions

An API capability to support management of application permissions for user administration.

There is no Apache Fineract functionality for creating or deleting permissions. Permissions come pre-installed.

Permissions are not updated, except in the case of enabling or disabling non-read transactions for Maker Checker functionality

## Roles

An API capability to support management of application roles for user administration.

## Users

An API capability to support administration of application users.

## AdhocQuery

## Maker Checker (or 4-eye) functionality

## default

## Run Reports

## Scheduler

## List Report Mailing Job History

## Fetch authenticated user details

## Mix Report

## Mix Taxonomy

## Mix Mapping

## Notification

## Provisioning Category

## Cashiers

## Cashier Journals

## Self Account transfer

## Self Third Party Transfer

## Self Client

## Self Loans

## Pocket

## Self Service Registration

## Self Savings Account

## Self User

## Self User Details

## Self Share Accounts

## Self Score Card

## Self Spm

## Self Dividend

## Score Card

## Spm-Surveys
