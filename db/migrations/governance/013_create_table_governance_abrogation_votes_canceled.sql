create table governance.abrogation_votes_canceled
(
    proposal_id       bigint  not null,
    user_id           text    not null,
    block_timestamp   bigint,
    tx_hash           text    not null,
    tx_index          integer not null,
    log_index         integer not null,
    logged_by         text    not null,
    included_in_block bigint  not null,
    created_at        timestamp default now()
);

create index abrogation_votes_canceled_idx
    on governance.abrogation_votes_canceled (proposal_id asc, user_id asc, block_timestamp desc);


---- create above / drop below ----

drop table if exists governance.gabrogation_votes_canceled;
drop index if exists governance.abrogation_votes_canceled_idx;