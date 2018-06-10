-- +migrate Up

-- +migrate StatementBegin
ALTER TABLE chainquery.claim
  DROP FOREIGN KEY claim_ibfk_2;
-- +migrate StatementEnd

-- +migrate StatementBegin
ALTER TABLE chainquery.input
  DROP FOREIGN KEY input_ibfk_1;
-- +migrate StatementEnd

-- +migrate StatementBegin
ALTER TABLE chainquery.support
  DROP FOREIGN KEY support_ibfk_1,
  CHANGE COLUMN transaction_hash transaction_hash_id VARCHAR(70) CHARACTER SET 'latin1' COLLATE 'latin1_general_ci' NULL DEFAULT NULL,
  ADD CONSTRAINT fk_transaction
    FOREIGN KEY (transaction_hash_id)
    REFERENCES chainquery.transaction (hash)
      ON DELETE CASCADE
      ON UPDATE NO ACTION;
-- +migrate StatementEnd

-- +migrate StatementBegin
ALTER TABLE chainquery.output
  DROP FOREIGN KEY output_ibfk_3,
  DROP FOREIGN KEY output_ibfk_2;
-- +migrate StatementEnd