module pi_go

go 1.15

require (
	pi_common v0.0.0
	pi_drive v0.0.0
    pi_server v0.0.0
    pi_client v0.0.0
)

replace (
    pi_common => ./pi_common
    pi_drive => ./pi_drive
    pi_server => ./pi_server
    pi_client => ./pi_client
 )
